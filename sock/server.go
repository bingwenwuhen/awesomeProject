package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	server := ":8330"
	addr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Println("resolve err:", err)
		return
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listen.Close()
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}
		go handle(accept)
	}
}

func handle(conn net.Conn) {
	fmt.Println("receive conn")
	go func() {
		for {
			var buf = make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					conn.Close()
					return
				}
				fmt.Println("read error: ", err)
				return
			}
			fmt.Println("read data: ", n, ":", string(buf))
		}
	}()

	curTime := time.Now().String()
	_, err := conn.Write([]byte(curTime))
	if err != nil {
		fmt.Println("write err：", err)
		return
	}
	fmt.Println("send data: ", curTime)
}
