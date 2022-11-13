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
	isEqual := true
	for i := 0; i < len(h.ip)-1; i++ {
		if h.ip[i] != h.ip[i+1] {
			isEqual = false
			break
		}
	}

	fmt.Printf(
		"%d ",
		h.number,
	)

	if isEqual {
		if h.ip[0] == "*" {
			for i := 0; i < len(h.ip); i++ {
				fmt.Printf("* ")
			}
		} else {
			fmt.Printf(" %s ", h.ip[0])
			for _, latencie := range h.latencies {
				fmt.Printf(" %v ", latencie)
			}
		}
		fmt.Print("\n")
		return
	}

	if h.ip[0] == "*" {
		fmt.Print(" *\n")
	} else {
		fmt.Printf(" %s %v\n", h.ip[0], h.latencies[0])
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
