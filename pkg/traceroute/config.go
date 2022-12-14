package traceroute

// Config is a configuration struct for traceroute.
type Config struct {
	Host          string
	FirstHop      int
	MaxHops       int
	TimeLimit     int64
	Port          int
	NumberQueries int
	Protocol      string
	PacketSize    int
	OnlyIP        bool
}
