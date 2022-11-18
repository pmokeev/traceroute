# Traceroute

#### *Traceroute* tracks the route packets take across an IP network on their way to a given host. It utilizes the IP protocol's time to live (TTL) field and attempts to elicit an **ICMP TIME_EXCEEDED** response from each gateway along the path to the host.

## Example

![Example](assets/example.gif)

## Available commands
```bash
$ go run cmd/main.go --help
  -M int
    	Set the initial time-to-live value used in outgoing probe packets. (default 1)
  -P string
    	Send packets of specified IP protocol. The currently supported protocols are: UDP (by default) and ICMP. (default "UDP")
  -d string
    	The required parameter is the name of the destination host.
  -m int
    	Specifies the maximum number of hops (max time-to-live value) traceroute will probe. (default 30)
  -p int
    	Port to connect. (default 33434)
  -ps int
    	Size of sending packet. (default 52)
  -q int
    	Set the number of probes per ttl to nqueries. (default 3)
  -w int
    	Set the time (in ms) to wait for a response to a probe. (default 50)
```