package main

import (
	"github.com/vojapet/golang-tplink-hs100/pkg/hs1x0"
	"time"
	"flag"
)

func main() {
	ip := flag.String("ip", "192.168.1.100", "Smart plug IP.")
	flag.Parse()

	h := hs1x0.NewHs110(*ip)

	println("Name of device:")
	name, _ := h.GetName()
	println(name)

	time.Sleep(2000 * time.Millisecond)

	println("Is on:")
	b, _ := h.IsOn()
	println(b)

	time.Sleep(2000 * time.Millisecond)

	println("Turning on")
	_ = h.TurnOn()
	println("done")

	time.Sleep(2000 * time.Millisecond)

	println("Is on:")
	b, _ = h.IsOn()
	println(b)

	time.Sleep(2000 * time.Millisecond)

	println("Turning off")
	_ = h.TurnOff()
	println("done")

	time.Sleep(2000 * time.Millisecond)

	println("Is on:")
	b, _ = h.IsOn()
	println(b)
}
