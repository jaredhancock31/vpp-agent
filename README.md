<h1 align="center">VPP Agent</h1>

- [CNTD Agent Quickstart](#cntd-agent-quickstart)
  - [Client example](#cl)

The VPP Agent is a Go implementation of a control/management plane for [VPP][vpp] based
cloud-native [Virtual Network Functions][vnf] (VNFs). The VPP Agent is built on top of 
[CN Infra][cn-infra], a framework for developing cloud-native VNFs (CNFs).

The VPP Agent can be used as-is as a management/control agent for VNFs  based on off-the-shelf
VPP (e.g. a VPP-based vswitch), or as a framework for developing management agents for VPP-based
CNFs. An example of a custom VPP-based CNF is the [Contiv-VPP][contiv-vpp] vswitch.

> Please note that the content of this repository is currently **WORK IN PROGRESS**!

## CNTD Agent Quickstart

CNTD Agent is a fork of VPP Agent. The primary difference is its use of an embedded BoltDB KV Store rather than
etcd. The dev Dockerfile has been modified to start a VPP instance as well as the CNTD Agent. Config files for bolt,
gRPC, VPP, and other plugins are in that same docker directory. 

To get started, install the necessary tools for building and whatnot, such as Go, Docker, etc...

```shell script
sudo apt-get update
sudo apt-get install build-essential
wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.14.linux-amd64.tar.gz
sudo chown -R $USER:$USER /usr/local/go
export PATH=$PATH:/usr/local/go/bin
sudo apt install make
```

Clone this repository, and build
```shell script
git clone https://github.com/jaredhancock31/vpp-agent.git
cd vpp-agent
make build
```

If you want to just rebuild the CNTD Agent image you can do
```shell script
# build both dev and prod
make images 

# dev only
make dev-image

# prod only 
make prod-image
```

Now, let's run the image. Note that we want our agent to configure the host network, so you'll make it `privileged` without
the typical docker bridge. 

```shell script
sudo docker run -it --name cntd-agent --privileged --network host --rm dev_vpp_agent
``` 

### Client example
> This is still a work-in-progress. The idea is two have two namespaces with veth pairs and a bridge domain. 

Note that the REST API is read-only, so we'll be sending configurations with a gRPC client, [found here](examples/cntd/main.go).

- In a separate terminal (still on the master):

    ```
    cd examples/cntd
    go run main.go
    ```

    Several flags can be set:
    * `-address=<address>` - for grpc server address/socket-file (otherwise localhost will be used)
    * `-socket-type=<type>` - options are tcp, tcp4, tcp6, unix or unixpacket. Defaults to tcp if not set
    * `request-period=<time_in_sec>` - time between grpc requests
    
    The example prints all received VPP notifications. 

- You can hit the REST API to get all interfaces with `curl -X GET http://localhost:9191/dump/vpp/v2/interfaces`. Obviously, you're 
 on the host already so you can just use things like `ethtool` and `ifconfig` as well. Whatever floats your boat.
 
- Delete interfaces with `sudo ip link del veth11`, etc.

## VPP Agent Quickstart

For a quick start with the VPP Agent, you can use the pre-built Docker images on DockerHub
that contain the VPP Agent and VPP: [ligato/vpp-agent][vpp-agent] (or for ARM64: [ligato/vpp-agent-arm64][vpp-agent-arm64]).

0. Start ETCD on your host (e.g. in Docker as described [here][etcd-local]).

   Note: for ARM64 see the information for [etcd][etcd-arm64].

1. Run VPP + VPP Agent in a Docker container:
```
docker run -it --rm --name agent1 --privileged ligato/vpp-agent
```

2. Manage VPP agent using agentctl:
```
docker exec -it agent1 agentctl --help
docker exec -it agent1 agentctl status
```

3. Check the configuration (via agentctl or in VPP console):
```
docker exec -it agent1 agentctl dump all
docker exec -it agent1 vppctl -s localhost:5002 show interface
```

**Next Steps**

See [README][docker-image] of development docker image for more details.

## Documentation
[![GoDoc](https://godoc.org/github.com/ligato/vpp-agent?status.svg)](https://godoc.org/github.com/ligato/vpp-agent)

Extensive documentation for the VPP Agent can be found at [docs.ligato.io](https://docs.ligato.io).

## Architecture

The VPP Agent is basically a set of VPP-specific plugins that use the 
CN-Infra framework to interact with other services/microservices in the
cloud (e.g. a KV data store, messaging, log warehouse, etc.). The VPP Agent
exposes VPP functionality to client apps via a higher-level model-driven 
API. Clients that consume this API may be either external (connecting to 
the VPP Agent via REST, gRPC API, Etcd or message bus transport), or local
Apps and/or Extension plugins running on the same CN-Infra framework in the 
same Linux process. 

The VNF Agent architecture is shown in the following figure: 

![vpp agent](docs/imgs/vpp_agent.png "VPP Agent & its Plugins on top of cn-infra")

Each (northbound) VPP API - L2, L3, ACL, ... - is implemented by a specific
VNF Agent plugin, which translates northbound API calls/operations into 
(southbound) low level VPP Binary API calls. Northbound APIs are defined 
using [protobufs][protobufs], which allow for the same functionality to be accessible
over multiple transport protocols (HTTP, gRPC, Etcd, ...). Plugins use the 
[GoVPP library][govpp] to interact with the VPP.

The following figure shows the VPP Agent in context of a cloud-native VNF, 
where the VNF's data plane is implemented using VPP/DPDK and 
its management/control planes are implemented using the VNF agent:

![context](docs/imgs/context.png "VPP Agent & its Plugins on top of cn-infra")

## Contributing

![GitHub contributors](https://img.shields.io/github/contributors/ligato/vpp-agent.svg)

If you are interested in contributing, please see the [contribution guidelines][contribution].

## License

[![GitHub license](https://img.shields.io/badge/license-Apache%20license%202.0-blue.svg)](https://github.com/ligato/vpp-agent/blob/master/LICENSE)

[agentctl]: cmd/agentctl
[cn-infra]: https://github.com/ligato/cn-infra
[contiv-vpp]: https://github.com/contiv/vpp
[contribution]: CONTRIBUTING.md
[docker]: docker
[docker-image]: https://docs.ligato.io/en/latest/user-guide/get-vpp-agent/#local-image-build
[etcd-arm64]: https://docs.ligato.io/en/latest/user-guide/arm64/#arm64-and-etcd-server
[etcd-local]: https://docs.ligato.io/en/latest/user-guide/get-vpp-agent/#connect-vpp-agent-to-the-key-value-data-store
[govpp]: https://wiki.fd.io/view/GoVPP
[ligato-docs]: http://docs.ligato.io/
[protobufs]: https://developers.google.com/protocol-buffers/
[vnf]: https://docs.ligato.io/en/latest/intro/glossary/#cnf
[vpp]: https://fd.io/vppproject/vpptech/
[vpp-agent]: https://hub.docker.com/r/ligato/vpp-agent
[vpp-agent-arm64]: https://hub.docker.com/r/ligato/vpp-agent-arm64
