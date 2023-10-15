package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	serverAddr := ":8330"
	addr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		fmt.Println("resolve err:", err)
		return
	}

	_, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}

	closed := make(chan bool)

	for {
		select {
		case <-closed:
			fmt.Println("服务端关闭")
			return
		}
	}
}
func handler(conn net.Conn) {
	go func() {
		for {
			var buf = make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					conn.Close()
					return
				}
				fmt.Println("read err:", err)
				return
			}
			fmt.Println("read data ", n, ":", string(buf))
		}
	}()

	curTime := time.Now().String()
	_, err := conn.Write([]byte(curTime))
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
	fmt.Println("send data", curTime)
}
