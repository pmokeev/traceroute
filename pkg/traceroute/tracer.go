package traceroute

import (
	"errors"
	"fmt"
	"net"
)

type Tracer struct {
	config *Config
	sender *sender
}

func NewTracer(config *Config) *Tracer {
	return &Tracer{
		config: config,
		sender: newSender(config),
	}
}

func (t *Tracer) resolveIP() (*string, error) {
	ips, err := net.LookupIP(t.config.Host)
	if err != nil {
		return nil, err
	}

	IPv4 := make([]string, 0, len(ips))
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			IPv4 = append(IPv4, ipv4.String())
		}
	}

	if len(IPv4) > 0 {
		if len(IPv4) > 1 {
			fmt.Printf(
				"traceroute: Warning: %s has multiple addresses; using %s\n",
				t.config.Host,
				IPv4[0],
			)
		}

		return &IPv4[0], nil
	}

	return nil, errors.New("no IPv4 address")
}

func (t *Tracer) convertIP(ip string) ([4]byte, error) {
	ipAddr, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		return [4]byte{}, err
	}

	destinationAddress := [4]byte{}
	copy(destinationAddress[:], ipAddr.IP.To4())
	return destinationAddress, nil
}

func (t *Tracer) Run() {
	destinationIP, err := t.resolveIP()
	if err != nil {
		fmt.Printf("traceroute: unknown host %s", t.config.Host)
		return
	}

	fmt.Printf(
		"traceroute to %s (%s), %d hops max, 52 byte packets\n",
		t.config.Host,
		*destinationIP,
		t.config.MaxHops,
	)

	convertedIP, err := t.convertIP(*destinationIP)
	if err != nil {
		return
	}

	TTL := t.config.FirstHop
	for {
		if TTL > t.config.MaxHops {
			return
		}

		currentHop := newHop(TTL, t.config.NumberQueries)

		for try := 0; try < t.config.NumberQueries; try++ {
			receivedMsg, err := t.sender.SendPacket(convertedIP, TTL)
			if err != nil {
				currentHop.insertRequest("*", "")
				continue
			}

			currentHop.insertRequest(receivedMsg.responceAddress, receivedMsg.latency)
		}

		currentHop.printHop()
		if currentHop.checkHop(*destinationIP) {
			return
		}
		TTL++
	}
}
