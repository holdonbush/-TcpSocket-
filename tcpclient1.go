package main

import (
	"os"
	"fmt"
	"net"
	"log"
	"io"
	"sync"
)

var value string
var wg sync.WaitGroup
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr,"Usage: %s host:post",os.Args[0])
		os.Exit(1)
	}
	addr := os.Args[1]

	conn, err := net.Dial("tcp",addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("访问的公网IP是：",conn.RemoteAddr().String())
	fmt.Println("客户端连接的地址及端口是：",conn.LocalAddr().String())
	buf := make([]byte,1024)
	wg.Add(2)
	go readcnt(conn,buf)
	go writecnt(conn)
	wg.Wait()
}

func writecnt(connect net.Conn)  {
	defer wg.Done()
	for {
		fmt.Scanln(&value)
		n, err := connect.Write([]byte(value))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("size: ",n)
	}

}

func readcnt(connect net.Conn,buf []byte) {
	defer wg.Done()
	for {
		length, err := connect.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		fmt.Println(string(buf[:length]))
		if value == "close" {
			connect.Close()
		}
	}
}