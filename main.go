package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	host := "127.0.0.1"
	port := "80"

	if len(os.Args) >= 2 {
		host = os.Args[1]
	}
	if len(os.Args) >= 3 {
		port = os.Args[2]
	}
	addrStr := host + ":" + port
	for {
		bt := time.Now()
		_, err := net.DialTimeout("tcp", addrStr, 5*time.Second)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		d := time.Now().Sub(bt)
		fmt.Printf("Ping %v time cost:%4dms\n", addrStr, d.Milliseconds())
		time.Sleep(time.Second)
	}
}
