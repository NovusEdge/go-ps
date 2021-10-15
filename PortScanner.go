package gops

import (
	"fmt"
	"net"
	"sync"
	"time"
)

//PortScanner : A port scanner
/*
	Domain   [string]
	Protocol [string]
*/
type PortScanner struct {
	Domain   string
	Protocol string
}

//Scan iteartes over all ports in range [startPort, endPort] and reports which ports are open by printing to stdout.
/*
	startPort [int]
	endPort   [int]
	timeout   [time.Duration]
*/
func (ps *PortScanner) Scan(startPort int, endPort int, timeout time.Duration) {
	var wg sync.WaitGroup

	for i := startPort; i <= endPort; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", ps.Domain, j)
			conn, err := net.DialTimeout(ps.Protocol, address, timeout)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("\033[32m[+] Port %d is Open.\033[0m\n", j)
		}(i)
	}
	wg.Wait()
}

//Scans iteartes over all ports in range [startPort, endPort] and returns a list of open ports.
/*
	startPort [int]
	endPort   [int]
	timeout   [time.Duration]
*/
func (ps *PortScanner) Scans(startPort int, endPort int, timeout time.Duration) (res []int) {
	var wg sync.WaitGroup
	var OpenPorts chan int

	for i := startPort; i <= endPort; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", ps.Domain, j)
			conn, err := net.DialTimeout(ps.Protocol, address, timeout)
			if err != nil {
				return
			}
			conn.Close()
			OpenPorts <- j
		}(i)
	}
	wg.Wait()

	close(OpenPorts)

	for i := 0; i < len(OpenPorts); i++ {
		res = append(res, i)
	}
	return
}
