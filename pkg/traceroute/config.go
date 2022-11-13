package traceroute

type Config struct {
	Host          string
	FirstHop      int
	MaxHops       int
	TimeLimit     int64
	Port          int
	NumberQueries int
	Protocol      string
	PacketSize    int
}
