package traceroute

import (
	"fmt"
)

type hop struct {
	number    int
	ip        []string
	latencies []string
}

func newHop(index int) *hop {
	return &hop{
		number:    index,
		ip:        make([]string, 0, 3),
		latencies: make([]string, 0, 3),
	}
}

func (h *hop) insertRequest(ip string, latency string) {
	h.ip = append(h.ip, ip)
	h.latencies = append(h.latencies, latency)
}

func (h *hop) printHop() {
	if h.ip[0] == h.ip[1] && h.ip[1] == h.ip[2] {
		if h.ip[0] == "*" {
			fmt.Printf(
				"%d * * *\n",
				h.number,
			)
		} else {
			fmt.Printf(
				"%d %s %v %v %v\n",
				h.number,
				h.ip[0],
				h.latencies[0],
				h.latencies[1],
				h.latencies[2],
			)
		}
		return
	}

	if h.ip[0] == "*" {
		fmt.Printf("%d *\n", h.number)
	} else {
		fmt.Printf("%d %s %v\n", h.number, h.ip[0], h.latencies[0])
	}

	for index := 1; index < len(h.ip); index++ {
		if h.ip[index] == "*" {
			fmt.Print("  *\n")
		} else {
			fmt.Printf("  %s %v\n", h.ip[index], h.latencies[index])
		}
	}
}

func (h *hop) checkHop(ip string) bool {
	return ip == h.ip[0] ||
		ip == h.ip[1] ||
		ip == h.ip[2]
}
