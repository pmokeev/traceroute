package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"

	"github.com/pmokeev/traceroute/pkg/traceroute"
)

func parseArgs() (*traceroute.Config, error) {
	host := flag.String("d", "ya.ru", "The only required parameter is the name of the destination host.")
	firstTTL := flag.String("m", "1", "Set the initial time-to-live value used in outgoing probe packets.")
	maxTTL := flag.String("M", "30", "Specifies the maximum number of hops (max time-to-live value) traceroute will probe.")
	waitTime := flag.String("w", "50", "Set the time (in ms) to wait for a response to a probe.")
	port := flag.String("p", "33434", "Port to connect.")
	protocol := flag.String("P", "UDP", "Send packets of specified IP protocol. The currently supported protocols are: UDP (by default) and ICMP.")
	nQueries := flag.String("q", "3", "Set the number of probes per ttl to nqueries.")
	packetSize := flag.String("ps", "52", "Size of sending packet.")
	flag.Parse()

	maxTTLConverted, err := strconv.Atoi(*maxTTL)
	if err != nil {
		return nil, err
	}

	firstTTLConverted, err := strconv.Atoi(*firstTTL)
	if err != nil {
		return nil, err
	}

	waitTimeConverted, err := strconv.Atoi(*waitTime)
	if err != nil {
		return nil, err
	}

	portConverted, err := strconv.Atoi(*port)
	if err != nil {
		return nil, err
	}

	nQueriesConverted, err := strconv.Atoi(*nQueries)
	if err != nil {
		return nil, err
	}

	if nQueriesConverted <= 0 {
		return nil, errors.New("number of probes must be positive")
	}

	packetSizeConverted, err := strconv.Atoi(*packetSize)
	if err != nil {
		return nil, err
	}

	switch *protocol {
	default:
		return nil, errors.New("unknown protocol")
	case "UDP", "ICMP":
	}

	return &traceroute.Config{
		Host:          *host,
		FirstHop:      firstTTLConverted,
		MaxHops:       maxTTLConverted,
		TimeLimit:     int64(waitTimeConverted),
		Port:          portConverted,
		NumberQueries: nQueriesConverted,
		Protocol:      *protocol,
		PacketSize:    packetSizeConverted,
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
