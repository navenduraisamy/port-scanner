package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type PortScanner struct {
	host    string
	port    int
	wg      *sync.WaitGroup
	timeout time.Duration
}

func (ps *PortScanner) Scan() {
	fmt.Printf("Scanning Host: %s\n", ps.host)
	if ps.port != 0 {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ps.host, ps.port), ps.timeout)
		if err != nil {
			return
		}
		fmt.Printf("Port: %v is open\n", ps.port)
		defer conn.Close()
		return
	}

	for port := 1; port <= 65535; port++ {
		ps.wg.Add(1)
		go func(p int) {
			defer ps.wg.Done()
			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ps.host, p), ps.timeout)
			if err != nil {
				return
			}
			fmt.Printf("Port: %v is open\n", p)
			defer conn.Close()
		}(port)
	}
	ps.wg.Wait()
}

func NewPortScanner(host string, port int) *PortScanner {
	return &PortScanner{
		host:    host,
		port:    port,
		wg:      &sync.WaitGroup{},
		timeout: time.Second,
	}
}
