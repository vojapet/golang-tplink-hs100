package main

import (
	"github.com/vojapet/golang-tplink-hs100/pkg/hs1x0"
	"log"
	"flag"
)

func main() {
	ipRange := flag.String("ip-range", "192.168.1.0/24", "Ip range for discovery.")
	flag.Parse()

	devices, err := hs1x0.Discover(*ipRange)

	if err != nil {
		panic(err)
	}

	log.Printf("Found devices: %d", len(devices))
	for _, d := range devices {
		name, _ := d.GetName()
		log.Printf("Device name: '%s', ip: '%s'", name, d.GetIp())
	}
}
