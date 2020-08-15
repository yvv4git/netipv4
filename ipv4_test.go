package netipv4

import (
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for GetIPv4AddressesFromNet
func TestGetIPv4AddressesFromNet(t *testing.T) {
	// сеточка
	ipNet := net.IPNet{
		IP:   net.IP{192, 168, 1, 0},
		Mask: net.IPv4Mask(255, 255, 255, 0),
	}

	ipV4AddressList, err := GetIPv4AddressesFromNet(&ipNet)
	if nil != err {
		panic(err)
	}

	assert.Equal(t, 255, len(ipV4AddressList))
}

// Test for GetIpv4FromIface
func TestGetIpv4FromIface(t *testing.T) {
	ifaceName := os.Getenv("TEST_IFACE")
	t.Log("Iface =", ifaceName)

	// get iface
	iface, err := net.InterfaceByName(ifaceName)
	if nil != err {
		panic("Don't find interface.")
	}

	ipV4Address := GetIpv4FromIface(iface)
	t.Log("Expected ip =", ipV4Address)

	assert.Equal(t, net.IP(net.IP{0xc0, 0xa8, 0x2, 0x2}), ipV4Address)

}

// Test for GetNetworkAddressFromIface
func TestGetNetworkAddressFromIface(t *testing.T) {
	ifaceName := os.Getenv("TEST_IFACE")
	t.Log("Iface =", ifaceName)

	// get iface
	iface, err := net.InterfaceByName(ifaceName)
	if nil != err {
		panic("Don't find interface.")
	}

	network := GetNetworkAddressFromIface(iface)
	t.Log(network)

	srcIPv4Address := os.Getenv("TEST_SRC_IP")
	assert.Equal(
		t,
		net.ParseIP(srcIPv4Address).To4(),
		network.IP.To4(),
	)
}
