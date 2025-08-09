package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"
)

func validateHostIP(address string) error {
	if net.ParseIP(address) == nil {
		return fmt.Errorf("invalid ip:%s", address)
	}
	return nil
}

func main() {

	hosts := flag.String("host", "127.0.0.1", "Target Hosts")
	port := flag.Int("port", 0, "Target port")
	flag.Parse()
	hostsList := strings.Split(*hosts, ",")

	start := time.Now()
	for _, host := range hostsList {
		if err := validateHostIP(host); err != nil {
			fmt.Println(err)
			continue
		}
		portScanner := NewPortScanner(strings.TrimSpace(host), *port)
		portScanner.Scan()
	}
	fmt.Println("Total time elapsed:", time.Since(start))

}
