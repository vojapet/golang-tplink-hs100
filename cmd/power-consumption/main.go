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

	p, err := h.GetCurrentPowerConsumption()
	if err != nil {
		log.Println("Error on accessing device")
		os.Exit(1)
	}

	log.Println("Current Power consumption:")
	log.Printf("Voltage: %d mV", p.Voltage)
	log.Printf("Current: %d mA", p.Current)
	log.Printf("Power: %d mW", p.Power)
}
