package main

import (
	"context"
	linux_namespace "go.ligato.io/vpp-agent/v3/proto/ligato/linux/namespace"
	vpp_l2 "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/l2"
	"log"
	"net"
	"sync"
	"time"

	"github.com/namsral/flag"
	"go.ligato.io/cn-infra/v2/agent"
	"go.ligato.io/cn-infra/v2/infra"
	"go.ligato.io/cn-infra/v2/logging/logrus"
	"google.golang.org/grpc"

	"go.ligato.io/vpp-agent/v3/proto/ligato/configurator"
	"go.ligato.io/vpp-agent/v3/proto/ligato/linux"
	linux_intf "go.ligato.io/vpp-agent/v3/proto/ligato/linux/interfaces"
	"go.ligato.io/vpp-agent/v3/proto/ligato/vpp"
	vpp_intf "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/interfaces"
)

// TODO use the const values in the functions
const (
	bdNetPrefix = "10.11.1."
	bdNetMask   = "/24"

	veth1LogicalName = "myVETH1"
	veth1HostName    = "veth1"
	veth1IPAddr      = bdNetPrefix + "1"
	veth1HwAddr      = "66:66:66:66:66:66"

	veth2LogicalName = "myVETH2"
	veth2HostName    = "veth2"

	afPacketLogicalName = "myAFPacket"
	afPacketHwAddr      = "a7:35:45:55:65:75"

	vppTapLogicalName = "myVPPTap"
	vppTapHwAddr      = "b3:12:12:45:A7:B7"
	vppTapVersion     = 2

	linuxTapLogicalName = "myLinuxTAP"
	linuxTapHostName    = "tap_to_vpp"
	linuxTapIPAddr      = bdNetPrefix + "2"
	linuxTapHwAddr      = "88:88:88:88:88:88"

	mycroservice1 = "microservice1"
	mycroservice2 = "microservice2"

	bviLoopName   = "myLoopback1"
	bviLoopIP     = bdNetPrefix + "3"
	bviLoopHwAddr = "cd:cd:cd:cd:cd:cd"

	loop2Name   = "myLoopback2"
	loop2HwAddr = "ef:ef:ef:ef:ef:ef"

	bdName                = "myBridgeDomain"
	bdFlood               = true
	bdUnknownUnicastFlood = true
	bdForward             = true
	bdLearn               = false /* Learning turned off, FIBs are needed for connectivity */
	bdArpTermination      = true
	bdMacAge              = 0
)

var exampleFinished = make(chan struct{})
var (
	timeout = flag.Int("timeout", 20, "Timeout between applying of initial and modified configuration in seconds")
	address    = flag.String("address", "172.17.0.1:9111", "address of GRPC server")
	socketType = flag.String("socket-type", "tcp", "socket type [tcp, tcp4, tcp6, unix, unixpacket]")
	cleanup = flag.Bool("cleanup", false, "cleanup all the stuff we made in the previous run")
	dialTimeout = time.Second * 2
)

/***********************************************
 * Here's what we want the end-result to be    *
 *                                             *
 *  +---------------------------------------+  *
 *  |       +-- Bridge domain --+           |  *
 *  |       |                   |           |  *
 *  | +-----+------+      +-----+------+    |  *
 *  | | afpacket1  |      | afpacket2  |    |  *
 *  | +-----+------+      +-----+------+    |  *
 *  |       |                   |           |  *
 *  +-------+-------------------+-----------+  *
 *          |                   |              *
 *  +-------+--------+  +-------+--------+     *
 *  | veth11         |  | veth21         |     *
 *  +-------+--------+  +-------+--------+     *
 *          |                   |              *
 *  +-------+---------+ +-------+---------+    *
 *  | veth12          | | veth22          |    *
 *  | IP: 10.0.0.1/24 | | IP: 10.0.0.2/24 |    *
 *  | NAMESPACE: ns1  | | NAMESPACE: ns2  |    *
 *  +-----------------+ +-----------------+    *
 ***********************************************/

func main()  {
	ep := &ExamplePlugin{}
	ep.SetName("cntd-client")
	ep.Setup()

	a := agent.NewAgent(
		agent.AllPlugins(ep),
		agent.QuitOnClose(exampleFinished),
	)
	if err := a.Run(); err != nil {
		log.Fatal()
	}

}

// ExamplePlugin demonstrates the use of the remoteclient to locally transport example configuration into the default VPP plugins.
type ExamplePlugin struct {
	infra.PluginDeps

	conn *grpc.ClientConn

	wg     sync.WaitGroup
	cancel context.CancelFunc
}

// Init initializes example plugin.
func (p *ExamplePlugin) Init() (err error) {
	// Set up connection to the server.
	p.conn, err = grpc.Dial("unix",
		grpc.WithInsecure(),
		grpc.WithDialer(dialer(*socketType, *address, dialTimeout)),
	)
	if err != nil {
		return err
	}

	client := configurator.NewConfiguratorServiceClient(p.conn)

	if *cleanup {
		p.Log.Info("Executing cleanup")
	} else {
		// Apply initial VPP configuration.
		go p.demonstrateClient(client)
	}

	// Schedule reconfiguration.
	var ctx context.Context
	ctx, p.cancel = context.WithCancel(context.Background())
	_ = ctx
	/*plugin.wg.Add(1)
	go plugin.reconfigureVPP(ctx)*/

	go func() {
		time.Sleep(time.Second * 30)
		close(exampleFinished)
	}()

	return nil
}

// Close cleans up the resources.
func (p *ExamplePlugin) Close() error {
	logrus.DefaultLogger().Info("Closing example plugin")

	p.cancel()
	p.wg.Wait()

	if err := p.conn.Close(); err != nil {
		return err
	}

	return nil
}

// Dialer for unix domain socket
func dialer(socket, address string, timeoutVal time.Duration) func(string, time.Duration) (net.Conn, error) {
	return func(addr string, timeout time.Duration) (net.Conn, error) {
		// Pass values
		addr, timeout = address, timeoutVal
		// Dial with timeout
		return net.DialTimeout(socket, addr, timeoutVal)
	}
}

// demonstrateClient propagates snapshot of the whole initial configuration to VPP plugins.
func (p *ExamplePlugin) demonstrateClient(client configurator.ConfiguratorServiceClient) {
	time.Sleep(time.Second * 2)
	p.Log.Infof("Requesting resync..")

	config := &configurator.Config{
		LinuxConfig: &linux.ConfigData{
			Interfaces: []*linux_intf.Interface{
				initialVeth11(), initialVeth12(),
			},
		},

		VppConfig: &vpp.ConfigData{
			Interfaces: []*vpp_intf.Interface{ afPacket1() },
		},
	}

	_, err := client.Update(context.Background(), &configurator.UpdateRequest{
		Update:     config,
		FullResync: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second * 5)
	p.Log.Infof("Requesting change..")

	_, err = client.Update(context.Background(), &configurator.UpdateRequest{
		Update: &configurator.Config{
			LinuxConfig: &linux.ConfigData{
				Interfaces: []*linux_intf.Interface{
					modifiedVeth11(), modifiedVeth12(), veth21(), veth22(),
				},
			},
			VppConfig: &vpp.ConfigData{
				Interfaces: []*vpp_intf.Interface{ afPacket2() },
			},
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

}

/**
TODO make a cleanup to get back to initial state
 */
//func (p *ExamplePlugin) tearDown(client configurator.ConfiguratorServiceClient)  {
//	time.Sleep(time.Second * 2)
//	p.Log.Infof("Requesting delete..")
//
//	ifaces = []*interfaces.Interface{v}
//	_, err = client.Delete(context.Background(), &configurator.DeleteRequest{
//		Delete: &configurator.Config{
//			VppConfig: &vpp.ConfigData{
//				Interfaces: initialVeth11(),
//			},
//		},
//	})
//}

/**
initial config
 */
func initialVeth11() *linux_intf.Interface {
	return &linux_intf.Interface{
		Name:    "veth11",
		Type:    linux_intf.Interface_VETH,
		Enabled: true,
		Link: &linux_intf.Interface_Veth{
			Veth: &linux_intf.VethLink{PeerIfName: "veth12"},
		},
	}
}

func initialVeth12() *linux_intf.Interface {
	return &linux_intf.Interface{
		Name:    "veth12",
		Type:    linux_intf.Interface_VETH,
		Enabled: true,
		Link: &linux_intf.Interface_Veth{
			Veth: &linux_intf.VethLink{PeerIfName: "veth11"},
		},
	}
}

func afPacket1() *vpp_intf.Interface {
	return &vpp_intf.Interface{
		Name:    "afpacket1",
		Type:    vpp_intf.Interface_AF_PACKET,
		Enabled: true,
		Link: &vpp_intf.Interface_Afpacket{
			Afpacket: &vpp_intf.AfpacketLink{
				HostIfName: "veth11",
			},
		},
	}
}

/**
modified config
 */
func modifiedVeth11() *linux_intf.Interface {
	return &linux_intf.Interface{
		Name:    "veth11",
		Type:    linux_intf.Interface_VETH,
		Enabled: true,
		Link: &linux_intf.Interface_Veth{
			Veth: &linux_intf.VethLink{PeerIfName: "veth12"},
		},
		Mtu: 1000,
	}
}

func modifiedVeth12() *linux_intf.Interface {
	return &linux_intf.Interface{
		Name:    "veth12",
		Type:    linux_intf.Interface_VETH,
		Enabled: true,
		Link: &linux_intf.Interface_Veth{
			Veth: &linux_intf.VethLink{PeerIfName: "veth11"},
		},
		IpAddresses: []string{"10.0.0.1/24"},
		PhysAddress: "D2:74:8C:12:67:D2",
		Namespace: &linux_namespace.NetNamespace{
			Reference: "ns1",
			Type:      linux_namespace.NetNamespace_NSID,
		},
	}
}

func afPacket2() *vpp_intf.Interface {
	return &vpp_intf.Interface{
		Name:    "afpacket2",
		Type:    vpp_intf.Interface_AF_PACKET,
		Enabled: true,
		Link: &vpp_intf.Interface_Afpacket{
			Afpacket: &vpp_intf.AfpacketLink{
				HostIfName: "veth21",
			},
		},
	}
}

func veth21() *linux_intf.Interface {
	return &linux_intf.Interface{
		Name:    "veth21",
		Type:    linux_intf.Interface_VETH,
		Enabled: true,
		Link: &linux_intf.Interface_Veth{
			Veth: &linux_intf.VethLink{PeerIfName: "veth22"},
		},
	}
}

func veth22() *linux_intf.Interface {
	return &linux_intf.Interface{
		Name:    "veth22",
		Type:    linux_intf.Interface_VETH,
		Enabled: true,
		Link: &linux_intf.Interface_Veth{
			Veth: &linux_intf.VethLink{PeerIfName: "veth21"},
		},
		IpAddresses: []string{"10.0.0.2/24"},
		PhysAddress: "92:C7:42:67:AB:CD",
		Namespace: &linux_namespace.NetNamespace{
			Reference: "ns2",
			Type:      linux_namespace.NetNamespace_NSID,
		},
	}
}


func bridgeDomain() *vpp_l2.BridgeDomain {
	return &vpp_l2.BridgeDomain{
		Name:                "br1",
		Flood:               true,
		UnknownUnicastFlood: true,
		Forward:             true,
		Learn:               true,
		ArpTermination:      false,
		MacAge:              0, /* means disable aging */
		Interfaces: []*vpp_l2.BridgeDomain_Interface{
			{
				Name:                    "afpacket1",
				BridgedVirtualInterface: false,
			}, {
				Name:                    "afpacket2",
				BridgedVirtualInterface: false,
			},
		},
	}
}
