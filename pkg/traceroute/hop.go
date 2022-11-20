package traceroute

import (
	"fmt"
	"net"
)

// hop is a struct for every hop in traceroute call.
type hop struct {
	number      int
	domainNames []string
	ips         []string
	latencies   []string
}

// newHop returns a new instance of hop.
func newHop(index, nQueries int) *hop {
	return &hop{
		number:      index,
		domainNames: make([]string, 0, nQueries),
		ips:         make([]string, 0, nQueries),
		latencies:   make([]string, 0, nQueries),
	}
}

// insertRequest inserts request into hop slices.
func (h *hop) insertRequest(ip, latency string) {
	resolvedHost := ""
	domainAddr, err := net.LookupAddr(ip)
	if err != nil {
		resolvedHost = ""
	} else {
		resolvedHost = domainAddr[0][:len(domainAddr[0])-1]
	}

	h.domainNames = append(h.domainNames, resolvedHost)
	h.ips = append(h.ips, ip)
	h.latencies = append(h.latencies, latency)
}

// printHop prints current hop into terminal.
func (h *hop) printHop() {
	isEqual := true
	for i := 0; i < len(h.ips)-1; i++ {
		if h.ips[i] != h.ips[i+1] {
			isEqual = false
			break
		}
	}

	fmt.Printf(
		"%d ",
		h.number,
	)

	if isEqual {
		if h.ips[0] == "*" {
			for i := 0; i < len(h.ips); i++ {
				fmt.Printf("* ")
			}
		} else {
			if h.domainNames[0] != "" {
				fmt.Printf(" %s (%s) ", h.domainNames[0], h.ips[0])
			} else {
				fmt.Printf(" %s ", h.ips[0])
			}
			for _, latencie := range h.latencies {
				fmt.Printf(" %v ", latencie)
			}
		}
		fmt.Print("\n")
		return
	}

	if h.ips[0] == "*" {
		fmt.Print(" *\n")
	} else {
		if h.domainNames[0] != "" {
			fmt.Printf(" %s (%s) %v\n", h.domainNames[0], h.ips[0], h.latencies[0])
		} else {
			fmt.Printf(" %s %v\n", h.ips[0], h.latencies[0])
		}
	}

	for index := 1; index < len(h.ips); index++ {
		if h.ips[index] == "*" {
			fmt.Print("  *\n")
		} else {
			if h.domainNames[index] != "" {
				fmt.Printf("  %s (%s) %v\n", h.domainNames[index], h.ips[index], h.latencies[index])
			} else {
				fmt.Printf("  %s %v\n", h.ips[index], h.latencies[index])
			}
		}
	}
}

// checkHop checks hop on reach some address.
func (h *hop) checkHop(ip string) bool {
	for _, hopIp := range h.ips {
		if ip == hopIp {
			return true
		}
	}

	return false
}
