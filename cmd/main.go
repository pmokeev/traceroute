package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/pmokeev/traceroute/pkg/traceroute"
)

func parseArgs() (*traceroute.Config, error) {
	host := flag.String("d", "ya.ru", "The only required parameter is the name of the destination host.")
	ttl := flag.String("m", "30", "Specifies the maximum number of hops (max time-to-live value) traceroute will probe. The default is 30.")
	waitTime := flag.String("w", "50", "Set the time (in ms) to wait for a response to a probe (default 50 ms).")
	port := flag.String("p", "33450", "Port to connect (default 33450).")
	flag.Parse()

	ttlConverted, err := strconv.Atoi(*ttl)
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

	return &traceroute.Config{
		Host:      *host,
		Hops:      ttlConverted,
		TimeLimit: int64(waitTimeConverted),
		Port:      portConverted,
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
