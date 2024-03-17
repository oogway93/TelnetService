package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleRemoteConn(conn net.Conn) {
	name := conn.RemoteAddr().String()
	fmt.Printf("%v connected\r\n", name)
	_, err := conn.Write([]byte("Hello, " + name + "\r\n"))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "Exit" {
			_, err := conn.Write([]byte("Bye\n"))
			if err != nil {
				panic(err)
			} else {
				fmt.Printf("%v disconnected\r\n", name)
				break
			}
		} else if text != "" {
			_, err := conn.Write([]byte("Your message: " + text + "\r\n"))
			if err != nil {
				panic(err)
			} else {
				fmt.Printf("Message from %v is %s\r\n", name, text)
			}
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleRemoteConn(conn)
	}
}
