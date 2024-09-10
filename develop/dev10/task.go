package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "Server conn timeout")
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Println("Right signature: dev10 [--timeout=10s] host port")
		return
	}
	address := net.JoinHostPort(flag.Arg(0), flag.Arg(1))

	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println("Server Closed")
		done <- struct{}{}
	}()

	go func() {
		io.Copy(conn, os.Stdin)
		done <- struct{}{}
	}()

	select {
	case <-interrupt:
		fmt.Println("Interrupted")
	case <-done:
		fmt.Println("Done")
	}
}
