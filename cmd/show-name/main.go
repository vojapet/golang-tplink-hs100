package main

import (
	"github.com/vojapet/golang-tplink-hs100/pkg/hs1x0"
	"log"
	"os"
	"flag"
)

func main() {
	ip := flag.String("ip", "192.168.1.100", "Smart plug IP.")
	flag.Parse()

	h := hs1x0.NewHs110(*ip)

	name, err := h.GetName()
	if err != nil {
		log.Print("Error on accessing device")
		os.Exit(1)
	}

	log.Printf("Name of device: %s", name)
}
