package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
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
	fmt.Printf("PING %v with tcp connecton\n", addrStr)

	succ := 0
	fail := 0

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		stat(addrStr, succ, fail)
		os.Exit(0)
	}()

	for {
		bt := time.Now()
		conn, err := net.DialTimeout("tcp", addrStr, 3*time.Second)
		d := time.Now().Sub(bt)
		if err != nil {
			fmt.Printf("ping %v error: %v, time cost:%4dms\n", addrStr, err, d.Milliseconds())
			fail++
		} else {
			fmt.Printf("ping %v (%v) time cost:%4dms\n", addrStr, conn.RemoteAddr(), d.Milliseconds())
			succ++
			conn.Close()
		}
		time.Sleep(time.Second)
	}
}

func stat(addrStr string, succ, fail int) {
	fmt.Printf("\n--- %v ping statistics ---\n", addrStr)
	fmt.Printf("%v requests send, %v responsed, %v lost, %v%% response lost\n", succ+fail, succ, fail, 100*fail/(succ+fail))
}
