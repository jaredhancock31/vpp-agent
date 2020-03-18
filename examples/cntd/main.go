package main

import (
	"context"
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

	dialTimeout = time.Second * 2
)
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

	// Apply initial VPP configuration.
	go p.demonstrateClient(client)

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
				initialVeth1(), initialVeth2(),
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
}

func initialVeth1() *linux_intf.Interface {
	return &linux_intf.Interface{
		Name:    "veth11",
		Type:    linux_intf.Interface_VETH,
		Enabled: true,
		Link: &linux_intf.Interface_Veth{
			Veth: &linux_intf.VethLink{PeerIfName: "veth12"},
		},
	}
}

func initialVeth2() *linux_intf.Interface {
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