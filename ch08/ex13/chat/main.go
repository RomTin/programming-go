// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel
type User struct {
	Client client
	Name   string
}

var (
	entering = make(chan User)
	leaving  = make(chan User)
	messages = make(chan string) // all incoming client messages
	users    = make(map[string]bool)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli.Client] = true
			users[cli.Name] = true

		case cli := <-leaving:
			delete(clients, cli.Client)
			delete(users, cli.Name)
			close(cli.Client)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	for u, _ := range users {
		ch <- "User " + u + " is in this room."
	}
	messages <- who + " has arrived"
	entering <- User{ch, who}
	timer := time.NewTimer(300 * time.Second)
	go func() {
		<-timer.C
		conn.Close()
	}()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(300 * time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- User{ch, who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
