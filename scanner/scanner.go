package scanner

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

const timeout = time.Millisecond * 500

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
			ipv6 = "[" + v.String() + "]"
			continue
		}
	}
	return
}

func GetOpenPorts(ip string, ports ...int) (open []int) {
	lock := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(len(ports))

	for _, v := range ports {
		go func(ip string, port int, lock sync.Locker, wg *sync.WaitGroup) {
			defer wg.Done()
			if isPortOpen(ip, port) {
				lock.Lock()
				open = append(open, port)
				lock.Unlock()
			}
		}(ip, v, lock, wg)
	}

	wg.Wait()
	return
}

func isPortOpen(ip string, port int) bool {
	target := fmt.Sprintf("%s:%d", ip, port)
	for {
		conn, err := net.DialTimeout("tcp", target, timeout)
		if err != nil {
			if strings.Contains(err.Error(), "too many open files") {
				time.Sleep(timeout)
				continue
			} else {
				return false
			}
		} else {
			conn.Close()
			return true
		}
	}
}
