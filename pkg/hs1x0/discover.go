package hs1x0

import (
	"github.com/pkg/errors"
	"net"
	"sync"
	"fmt"
)

type Barrier struct {
	wall chan int
}

func NewBarrier(count int) *Barrier {
	new_barrier := Barrier {
		wall: make(chan int, count),
	}
	return &new_barrier
}

func (self *Barrier) enter() {
	self.wall <- 1
}

func (self *Barrier) exit() {
	<- self.wall
}


func Discover(subnet string) ([]*Hs110, error) {
	ips, err := getIpAddresses(subnet)
	if err != nil {
		return nil, err
	}

	result := &discoverResult{
		devices: make([]*Hs110, 0),
		Mutex:   sync.Mutex{},
	}

	b := NewBarrier (10)	
	var wg sync.WaitGroup
	wg.Add(len(ips))
	for _, current := range ips {
		go tryIp(result, &wg, current, b)
	}
	wg.Wait()

	fmt.Println("")
	
	return result.devices, nil
}

func tryIp(r *discoverResult, wg *sync.WaitGroup, ip string, b *Barrier) {
	defer wg.Done()
	b.enter()
	defer b.exit()

	fmt.Print(".")
	
	hs110 := NewHs110(ip)
	_, err := hs110.GetName()
	if err != nil {
		return
	}

	r.Lock()
	defer r.Unlock()

	r.devices = append(r.devices, hs110)
}

type discoverResult struct {
	devices []*Hs110
	sync.Mutex
}

func getIpAddresses(subnet string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, errors.Wrap(err, "invalid subnet specfied")
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips[1 : len(ips)-1], nil
}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j > 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
