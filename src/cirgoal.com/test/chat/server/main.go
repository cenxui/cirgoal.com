package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func main() {

	listen, err := net.Listen("tcp", "localhost:8000");

	if err != nil {
		log.Fatal("error")
	}

	go broadcast()

	for {
		conn, err := listen.Accept()

		if err != nil{
			log.Print("connt error")
			continue
		}

		go handleRequest(conn)
	}

}

type client chan <- string

var (
	enter = make(chan client)
	leave = make(chan client)
	message = make(chan string)
)

func handleRequest(conn net.Conn)  {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who :=  conn.RemoteAddr().String()

	ch <- "You are " + who
	message <- who + " is entering"
	enter <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + ": " + input.Text() + "\n"
	}

	leave <- ch
	message <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch chan string)  {
	for msg := range ch {
		fmt.Fprintf(conn, msg);
	}
}

func broadcast()  {
	clients := make(map[client] bool)

	for {
		select {
		case s := <- message:
			for cli := range clients {
				cli <- s
			}

		case cli := <- enter:
			clients[cli] = true

		case cli := <- leave:
			delete(clients,cli);
			close(cli)
		}

	}
}