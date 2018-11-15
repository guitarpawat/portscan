package scanner

import (
	"net"
)

func NSLookUp(host string) (ipv4, ipv6 string, err error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return
	}
	for _, v := range ips {
		if len(ipv4) != 0 && len(ipv6) != 0 {
			break
		}
		v4 := v.To4()
		if v4 != nil && len(ipv4) == 0 {
			ipv4 = v.String()
			continue
		}
		v6 := v.To16()
		if v6 != nil && len(ipv6) == 0 {
			ipv6 = v.String()
			continue
		}
	}
	return
}

func GetOpenPorts(ip string, ports ...int) (open []int, err error) {
	return
}
