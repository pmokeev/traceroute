package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/pmokeev/traceroute/pkg/traceroute"
)

func parseArgs() (*traceroute.Config, error) {
	host := flag.String("d", "", "The only required parameter is the name of the destination host.")
	firstTTL := flag.Int("M", 1, "Set the initial time-to-live value used in outgoing probe packets.")
	maxTTL := flag.Int("m", 30, "Specifies the maximum number of hops (max time-to-live value) traceroute will probe.")
	waitTime := flag.Int64("w", 50, "Set the time (in ms) to wait for a response to a probe.")
	port := flag.Int("p", 33434, "Port to connect.")
	protocol := flag.String("P", "UDP", "Send packets of specified IP protocol. The currently supported protocols are: UDP (by default) and ICMP.")
	nQueries := flag.Int("q", 3, "Set the number of probes per ttl to nqueries.")
	packetSize := flag.Int("ps", 52, "Size of sending packet.")
	flag.Parse()

	if *nQueries <= 0 {
		return nil, errors.New("number of probes must be positive")
	}

	switch *protocol {
	default:
		return nil, errors.New("unknown protocol")
	case "UDP", "ICMP":
	}

	return &traceroute.Config{
		Host:          *host,
		FirstHop:      *firstTTL,
		MaxHops:       *maxTTL,
		TimeLimit:     *waitTime,
		Port:          *port,
		NumberQueries: *nQueries,
		Protocol:      *protocol,
		PacketSize:    *packetSize,
	}, nil
}

func main() {
	tracerouteConfig, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	tracer := traceroute.NewTracer(tracerouteConfig)
	tracer.Run()
}
