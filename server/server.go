package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	// "time"
)

var clients = make([]net.Conn, 0)

func deleteClient(client net.Conn) {
	var index int
	for i := 0; i < len(clients); i += 1 {
		if clients[i] == client {
			index = i 
		}
	}
	clients[index] = clients[len(clients) - 1]
	clients = clients[:len(clients) - 1]
}

func broadcast(addr, msg string) {
	for i := 0; i < len(clients); i += 1 {
		// fmt.Print(clients[i].RemoteAddr().String() + " " + addr)
		if clients[i].RemoteAddr().String() == addr {
			myMsg := clients[i].RemoteAddr().String() + "\n"
			clients[i].Write([]byte(myMsg))
			continue
		} 
		clients[i].Write([]byte(msg))
	}
}

func handleConnection(c net.Conn) {
	msg := c.RemoteAddr().String() + " entered the room.\n"
	broadcast(c.RemoteAddr().String(), msg)
	clients = append(clients, c)
	fmt.Print(msg)
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			deleteClient(c)
			return
		}
		if strings.TrimSpace(strings.ToUpper(string(netData))) == "EXIT" {
			msg = c.RemoteAddr().String() + " left the room.\n"
			broadcast(c.RemoteAddr().String(), msg)
			fmt.Print(msg)
			deleteClient(c)
			break
		}

		msg = c.RemoteAddr().String() + " -> " + string(netData)
		broadcast(c.RemoteAddr().String(), msg)
		fmt.Print(msg)
		// t := time.Now()
		// myTime := t.Format(time.RFC3339) + "\n"
		// blank := "\n"
		// c.Write([]byte(blank))
	}
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		// clients = append(clients, c)
		// fmt.Print(len(clients))
		go handleConnection(c)
	}
}
