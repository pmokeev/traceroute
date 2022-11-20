package main

import (
	"fmt"

	"github.com/pmokeev/traceroute/pkg/traceroute"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	host          = kingpin.Arg("host", "The required parameter is the name of the destination host.").Required().String()
	minTTL        = kingpin.Flag("first_ttl", "Set the initial time-to-live value used in outgoing probe packets.").Short('M').Default("1").Int()
	maxTTL        = kingpin.Flag("max_ttl", "Specifies the maximum number of hops (max time-to-live value) traceroute will probe.").Short('m').Default("30").Int()
	timeLimit     = kingpin.Flag("wait", "Set the time (in ms) to wait for a response to a probe.").Short('w').Default("50").Int64()
	port          = kingpin.Flag("port", "Port to connect.").Short('p').Default("33434").Int()
	onlyIP        = kingpin.Flag("onlyip", "Print hop addresses numerically rather than symbolically and numerically (saves a nameserver address-to-name lookup for each gateway found on the path).").Short('n').Default("false").Bool()
	numberQueries = kingpin.Flag("nqueries", "Set the number of probes per ttl to nqueries.").Short('q').Default("3").Int()
	protocol      = kingpin.Flag("proto", "Send packets of specified IP protocol. The currently supported protocols are: UDP (by default) and ICMP.").Short('P').Default("UDP").String()
	packetSize    = kingpin.Arg("packetSize", "Size of sending packet.").Default("50").Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	if *numberQueries <= 2 {
		fmt.Println("numberQueries must be >= 3")
		return
	}

	switch *protocol {
	case "UDP", "ICMP":
	default:
		fmt.Println("unsupported type of protocol")
		return
	}

	tracerouteConfig := &traceroute.Config{
		Host:          *host,
		FirstHop:      *minTTL,
		MaxHops:       *maxTTL,
		TimeLimit:     *timeLimit,
		Port:          *port,
		NumberQueries: *numberQueries,
		Protocol:      *protocol,
		PacketSize:    *packetSize,
		OnlyIP:        *onlyIP,
	}

	tracer := traceroute.NewTracer(tracerouteConfig)
	tracer.Run()
}
