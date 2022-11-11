package traceroute

import (
	"errors"
	"fmt"
	"net"

	"golang.org/x/net/ipv4"
)

type hopBlock struct {
	ip        string
	latencies []string
}

func (hb *hopBlock) printBlock(index int) {
	fmt.Printf("%d %s ", index, hb.ip)

	for _, currentTime := range hb.latencies {
		if currentTime == "*" {
			fmt.Print("* ")
		} else {
			fmt.Printf("%v ", currentTime)
		}
	}

	fmt.Printf("\n")
}

type Tracer struct {
	config *Config
}

func NewTracer(config *Config) *Tracer {
	return &Tracer{
		config: config,
	}
}

func (t *Tracer) resolveIP() (*string, error) {
	ips, err := net.LookupIP(t.config.Host)
	if err != nil {
		return nil, err
	}

	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			ipv4Converted := ipv4.String()
			return &ipv4Converted, nil
		}
	}

	return nil, errors.New("no IPv4 address")
}

func (t *Tracer) Run() {
	destinationIP, err := t.resolveIP()
	if err != nil {
		fmt.Printf("traceroute: unknown host %s", t.config.Host)
		return
	}

	fmt.Printf(
		"traceroute to %s (%s), %d hops max, %d byte packets\n",
		t.config.Host,
		*destinationIP,
		t.config.Hops,
		t.config.PacketSize,
	)

	TTL := 1
	for {
		if TTL > t.config.Hops {
			// TODO:
			return
		}

		currentHop := &hopBlock{
			latencies: make([]string, 3),
		}

		for try := 0; try < 3; try++ {
			receivedMsg, err := t.sendICMPPacket(*destinationIP, TTL)
			if err != nil {
				currentHop.latencies[try] = "*"
				continue
			}

			switch receivedMsg.responseCode {
			case int(ipv4.ICMPTypeEchoReply):
				currentHop.ip = receivedMsg.responceAddress
				currentHop.latencies[try] = receivedMsg.latency.String()
			default:
				return
			}
		}

		currentHop.printBlock(TTL)
		if currentHop.ip == *destinationIP {
			return
		}

		TTL++
	}
}
