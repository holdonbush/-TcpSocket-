package main

import (
	"net"
	"log"
	"io"
	"fmt"
)

var globalRoom []net.Conn

func main() {
	addr := "0.0.0.0:8080"
	listener, err := net.Listen("tcp",addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	buf := make([]byte,1024)
	globalRoom = append(globalRoom, conn)
	conn.Write([]byte("输入: "))
	for {
		length, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		fmt.Println(string(buf[:length]))
		for _,cnn := range globalRoom {
			value := string(buf[:length])
			cnn.Write([]byte("\n用户群聊消息："+value+"\n输入："))
		}
		if string(buf[:length]) == "close" {
			for index, cnn := range globalRoom {
				if cnn == conn {
					globalRoom = append(globalRoom[:index],globalRoom[index+1:]...)
					break
				}
			}
			break
		}
	}
	conn.Close()
}