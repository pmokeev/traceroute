# Traceroute

*Traceroute* tracks the route packets take across an IP network on their way to a given host. It utilizes the IP protocol's time to live (TTL) field and attempts to elicit an **ICMP TIME_EXCEEDED** response from each gateway along the path to the host.

## Example

![Example](assets/example.gif)

Before running the application, you must install [Golang 1.19+](https://go.dev/dl/) and and enter the following command:

```bash
$ go get ./...
```

**Example of usage:**
```bash
$ sudo go run cmd/main.go google.com
```

## Available commands

If you want to get the full list of available flags, you can get them by using the flag `--help`

```bash
$ go run cmd/main.go --help
usage: main [<flags>] <host> [<packetSize>]

Flags:
      --help         Show context-sensitive help (also try --help-long and --help-man).
  -M, --first_ttl=1  Set the initial time-to-live value used in outgoing probe packets.
  -m, --max_ttl=30   Specifies the maximum number of hops (max time-to-live value) traceroute will probe.
  -w, --wait=50      Set the time (in ms) to wait for a response to a probe.
  -p, --port=33434   Port to connect.
  -n, --onlyip       Print hop addresses numerically rather than symbolically and numerically (saves a nameserver address-to-name lookup for each gateway
                     found on the path).
  -q, --nqueries=3   Set the number of probes per ttl to nqueries.
  -P, --proto="UDP"  Send packets of specified IP protocol. The currently supported protocols are: UDP (by default) and ICMP.
      --version      Show application version.

Args:
  <host>          The required parameter is the name of the destination host.
  [<packetSize>]  Size of sending packet.
```