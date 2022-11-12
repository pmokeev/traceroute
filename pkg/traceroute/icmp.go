package traceroute

import (
	"fmt"
	"syscall"
	"time"
)

type icmpResponse struct {
	responceAddress string
	latency         string
}

func (t *Tracer) sendICMPPacket(destinationIP [4]byte, ttl int) (*icmpResponse, error) {
	receiveSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		return nil, err
	}
	defer syscall.Close(receiveSocket)

	latency := syscall.NsecToTimeval(1000 * 1000 * t.config.TimeLimit)
	syscall.SetsockoptTimeval(receiveSocket, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &latency)

	syscall.Bind(receiveSocket,
		&syscall.SockaddrInet4{
			Port: t.config.Port,
			Addr: [4]byte{0, 0, 0, 0},
		},
	)

	sendSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		return nil, err
	}
	defer syscall.Close(sendSocket)

	syscall.SetsockoptInt(sendSocket, 0x0, syscall.IP_TTL, ttl)
	start := time.Now()
	syscall.Sendto(sendSocket, make([]byte, 52), 0,
		&syscall.SockaddrInet4{
			Port: t.config.Port,
			Addr: destinationIP,
		},
	)

	receivedPacket := make([]byte, 1500) // MTU
	_, routerIP, err := syscall.Recvfrom(receiveSocket, receivedPacket, 0)

	if err == nil {
		responseAddress := routerIP.(*syscall.SockaddrInet4).Addr

		return &icmpResponse{
			latency: time.Since(start).String(),
			responceAddress: fmt.Sprintf(
				"%v.%v.%v.%v",
				responseAddress[0],
				responseAddress[1],
				responseAddress[2],
				responseAddress[3],
			),
		}, nil
	}

	return &icmpResponse{
		responceAddress: "*",
	}, err
}
