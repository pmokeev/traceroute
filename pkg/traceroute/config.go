package traceroute

type Config struct {
	Host      string
	Hops      int
	TimeLimit int64
	Port      int
}
