package traceroute

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// sender is a struct for sends ICMP, UDP and etc protocols packets.
type sender struct {
	config *Config
}

// newSender returns a new instance of sender.
func newSender(config *Config) *sender {
	return &sender{
		config: config,
	}
}

// createPacket creates some protocol packet.
func (s *sender) createPacket() ([]byte, error) {
	switch s.config.Protocol {
	case "ICMP":
		msg := &icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Seq:  0,
				Data: make([]byte, s.config.PacketSize-8),
			},
		}

		return msg.Marshal(nil)
	case "UDP":
		return make([]byte, s.config.PacketSize), nil
	}

	return nil, errors.New("unsupported type of protocol")
}

// SendPacket sends request by IP request using some protocol.
func (s *sender) SendPacket(destinationIP [4]byte, ttl int) (*response, error) {
	receiveSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		return nil, err
	}
	defer syscall.Close(receiveSocket)

	latency := syscall.NsecToTimeval(1000 * 1000 * s.config.TimeLimit)
	syscall.SetsockoptTimeval(receiveSocket, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &latency)

	if err := syscall.Bind(receiveSocket,
		&syscall.SockaddrInet4{
			Addr: [4]byte{0, 0, 0, 0},
		},
	); err != nil {
		return nil, err
	}

	var sendSocket int
	defer syscall.Close(sendSocket)

	switch s.config.Protocol {
	case "UDP":
		sendSocket, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
		if err != nil {
			return nil, err
		}
	case "ICMP":
		sendSocket, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_ICMP)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported type of protocol")
	}

	protocolMsg, err := s.createPacket()
	if err != nil {
		return nil, err
	}

	syscall.SetsockoptInt(sendSocket, 0x0, syscall.IP_TTL, ttl)
	start := time.Now()

	switch s.config.Protocol {
	case "UDP":
		if err := syscall.Sendto(sendSocket, protocolMsg, 0,
			&syscall.SockaddrInet4{
				Port: s.config.Port,
				Addr: destinationIP,
			},
		); err != nil {
			return nil, err
		}
	case "ICMP":
		if err := syscall.Sendto(sendSocket, protocolMsg, 0,
			&syscall.SockaddrInet4{
				Addr: destinationIP,
			},
		); err != nil {
			return nil, err
		}
	}

	receivedPacket := make([]byte, 1500) // MTU
	_, routerIP, err := syscall.Recvfrom(receiveSocket, receivedPacket, 0)

	if err == nil {
		responseAddress := routerIP.(*syscall.SockaddrInet4).Addr

		return &response{
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

	return &response{
		responceAddress: "*",
	}, err
}
