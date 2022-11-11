package main

import "github.com/pmokeev/traceroute/pkg/traceroute"

func main() {
	tracerConfig := &traceroute.Config{
		Host:       "ya.ru",
		Hops:       30,
		PacketSize: 52,
		TimeLimit:  1000,
	}

	tracer := traceroute.NewTracer(tracerConfig)

	tracer.Run()
}
