package netutils

import (
	"github.com/gooderbrother/nmap-IPrange"
)

func Cidr2IPs(cidrInp string) []string {
	res := []string{}
	res, _ = nmapIPrange.Handler(cidrInp)
	return res
}

//	func Cidr2IPs(cidr string) []string {
//		var addrs []string
//		block := a.NewIPAddressString(cidr).GetAddress()
//		for i := block.Iterator(); i.HasNext(); {
//			addrs = append(addrs, i.Next().String())
//		}
//		return addrs
//	}
//func Cidr2IPs(cidr string) []string {
//	ip, ipnet, err := net.ParseCIDR(cidr)
//	if err != nil {
//		return nil
//	}
//
//	var ips []string
//	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
//		ips = append(ips, ip.String())
//	}
//
//	// remove network address and broadcast address
//	lenIPs := len(ips)
//	switch {
//	case lenIPs < 2:
//		return ips
//
//	default:
//		return ips[1 : len(ips)-1]
//	}
//}
//
//func inc(ip net.IP) {
//	for j := len(ip) - 1; j >= 0; j-- {
//		ip[j]++
//		if ip[j] > 0 {
//			break
//		}
//	}
//}

//
//func Cidr2IPs(cidr string) []string {
//	// C段转ip
//	var ips []string
//
//	ipAddr, ipNet, err := net.ParseCIDR(cidr)
//	if err != nil {
//		log.Print(err)
//	}
//
//	for ip := ipAddr.Mask(ipNet.Mask); ipNet.Contains(ip); increment(ip) {
//		ips = append(ips, ip.String())
//	}
//
//	// CIDR too small eg. /31
//	if len(ips) <= 2 {
//		log.Print("err")
//	}
//
//	return ips
//}
//func increment(ip net.IP) {
//	for i := len(ip) - 1; i >= 0; i-- {
//		ip[i]++
//		if ip[i] != 0 {
//			break
//		}
//	}
//}
