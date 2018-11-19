// Package scanner provides useful functions for scanning the open port.
package scanner

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

const timeout = time.Millisecond * 500
const limit = 4

// NSLookUp provides a IP lookup for specified host,
// returns IPv4 and IPv6 strings (if have).
func NSLookUp(host string) (ipv4 string, err error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return
	}
	for _, v := range ips {
		v4 := v.To4()
		if v4 != nil && len(ipv4) == 0 {
			ipv4 = v.String()
			break
		}
	}
	return
}

// GetOpenPorts find the open port on specified host
// with the slice of ports of host provided,
// returns the slice of open ports of the host.
func GetOpenPorts(ip string, ports ...int) (open []int) {
	lock := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(len(ports))
	// A channel with buffer length limit
	sem := make(chan struct{}, limit)
	for _, port := range ports {
		// Block until the channel buffer available
		sem <- struct{}{}
		go func(ip string, port int, wg *sync.WaitGroup, lock sync.Locker, sem <- chan struct{}) {
			defer wg.Done()
			// Free channel buffer
			defer func() {<- sem}()

			if isPortOpen(ip, port) {
				lock.Lock()
				open = append(open, port)
				lock.Unlock()
			}
		}(ip, port, wg, lock, sem)
	}

	wg.Wait()
	close(sem)
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
