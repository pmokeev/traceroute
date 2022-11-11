package traceroute

import (
	"net"
	"os"
	"strings"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type icmpResponse struct {
	responseCode    int
	responceAddress string
	latency         time.Duration
}

func (t *Tracer) createICMPPacket() ([]byte, error) {
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte(""),
		},
	}

	return msg.Marshal(nil)
}

func (t *Tracer) sendICMPPacket(ip string, ttl int) (*icmpResponse, error) {
	icmpPacket, err := t.createICMPPacket()
	if err != nil {
		return nil, err
	}

	connection, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		return nil, err
	}
	defer connection.Close()

	if err := connection.IPv4PacketConn().SetTTL(ttl); err != nil {
		return nil, err
	}

	udpAddress := &net.UDPAddr{
		IP: net.ParseIP(ip),
	}

	start := time.Now()
	if _, err := connection.WriteTo(icmpPacket, udpAddress); err != nil {
		return nil, err
	}

	if err = connection.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(t.config.TimeLimit))); err != nil {
		return nil, err
	}

	reply := make([]byte, 1500)
	replySize, respondeAddr, err := connection.ReadFrom(reply)
	if err != nil {
		return nil, err
	}

	icmpMessage, err := icmp.ParseMessage(1, reply[:replySize])
	if err != nil {
		return nil, err
	}

	return &icmpResponse{
		latency:         time.Since(start),
		responseCode:    icmpMessage.Code,
		responceAddress: strings.Split(respondeAddr.String(), ":")[0],
	}, nil
}
