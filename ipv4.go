package netipv4

import (
	"encoding/binary"
	"errors"
	"net"
)

// Get IPv4 list from IPNET
func GetIPv4AddressesFromNet(netAddr *net.IPNet) (out []net.IP, err error) {
	var ipv4 net.IP

	if ipv4 = netAddr.IP.To4(); ipv4 == nil {
		err = errors.New("It's not ip version 4 address")
		return nil, err
	}

	num := binary.BigEndian.Uint32([]byte(ipv4))
	mask := binary.BigEndian.Uint32([]byte(netAddr.Mask))
	num &= mask
	for mask < 0xffffffff {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], num)
		out = append(out, net.IP(buf[:]))
		mask++
		num++
	}
	return
}

// Get IPv4 from network interface
func GetIpv4FromIface(iface *net.Interface) (ipv4Address net.IP) {
	cidrList, err := iface.Addrs()
	if nil != err {
		panic("Don't found ip addresses")
	}

	for _, cidr := range cidrList {
		if ipnet, ok := cidr.(*net.IPNet); ok {
			if ipv4 := ipnet.IP.To4(); ipv4 != nil {
				ipv4Address = ipv4
				break
			}
		}
	}

	return
}

// Get IPNet from network interface
func GetNetworkAddressFromIface(iface *net.Interface) (netAddress *net.IPNet) {
	cidrList, err := iface.Addrs()
	if nil != err {
		panic("Don't found ip addresses")
	}

	for _, cidr := range cidrList {
		if ipnet, ok := cidr.(*net.IPNet); ok {
			if ipv4 := ipnet.IP.To4(); ipv4 != nil {
				netAddress = ipnet
				break
			}
		}
	}

	return
}
