package ping

import (
	"fmt"
	"net"
	"time"
)

func Run(host string) error {

	ips, err := net.LookupIP(host)

	var ipv4 string
	var port = "443"

	if err != nil {
		return err
	}

	for _, ip := range ips {
		if ip.To4() != nil {
			ipv4 = ip.String()
			break

		}
	}

	address := net.JoinHostPort(ipv4, port)

	start := time.Now()
	conn, err := net.Dial("tcp", address)

	if err != nil {
		fmt.Println("Connection Failed to the address: " + address)
		return nil
	}
	defer conn.Close()
	elapsed := time.Since(start).Milliseconds()

	fmt.Println("Connection successful to address: " + address)
	fmt.Printf("The total time was: %d ms\n", elapsed)

	return nil

}
