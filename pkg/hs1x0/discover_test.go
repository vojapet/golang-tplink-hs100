package hs1x0

import (
	"testing"
	"net"
)

func Test_getIpAddresses_full24range(t *testing.T) {
	expectedNumberOfIps := 254
	
	testNet := "192.168.0.1/24"

	ipList, err := getIpAddresses(testNet)

	if err != nil {
		t.Errorf("%s", err)
	}

	if len(ipList) != expectedNumberOfIps {
		t.Errorf("Expected number of IPs [%d] received [%d].", expectedNumberOfIps, len(ipList))
	}
}

func Test_getIpAddresses_OneIP(t *testing.T) {
	expectedNumberOfIps := 1
	expectedIP := "192.168.0.100"
	
	testNet := "192.168.0.100/32"

	ipList, err := getIpAddresses(testNet)

	t.Log(ipList)
	
	if err != nil {
		t.Errorf("%s", err)
	}

	if len(ipList) != expectedNumberOfIps {
		t.Errorf("Expected number of IPs [%d] received [%d].", expectedNumberOfIps, len(ipList))
	}

	if ipList[0] != expectedIP {
		t.Errorf("Expected IP [%s] received [%s].", ipList[0], expectedIP)		
	}
}

func Test_inc_Basic (t *testing.T) {
	expectedIP := net.ParseIP("192.168.1.11")
	testIP := net.ParseIP("192.168.1.10")
	inc(testIP)
	
	if ! testIP.Equal(expectedIP) {
		t.Errorf("Expected [%s] but received [%s].", expectedIP, testIP)
	}
	
}

func Test_inc_Over24(t *testing.T) {
	expectedIP := net.ParseIP("192.168.2.0")
	testIP := net.ParseIP("192.168.1.255")
	inc(testIP)
	
	if ! testIP.Equal(expectedIP) {
		t.Errorf("Expected [%s] but received [%s].", expectedIP, testIP)
	}
	
}
