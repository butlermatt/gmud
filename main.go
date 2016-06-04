package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/butlermatt/gmud/client"
	"github.com/butlermatt/gmud/command"
)

func main() {
	var port = flag.Int("port", 4444, "port number")
	flag.Parse()
	var p = fmt.Sprintf(":%d", *port)

	Run(p)
}

// Run will initialize and start the server.
func Run(addr string) {
	srv, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		log.Println("Server shutting down.")
		err := srv.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	clients := make(map[string]*client.Client)

	msg := make(chan command.Commander)
	defer close(msg)
	add := make(chan *client.Client)
	defer close(add)
	rm := make(chan *client.Client)
	defer close(rm)

	go acceptConnections(srv, rm, add, msg)

	var shutdown = false
	for {
		select {
		case cmd := <-msg:
			if cmd.Log() {
				log.Printf("%s used %s", cmd.Player().Name(), cmd.Name())
			}
			if cmd.Name() == "shutdown" {
				for _, conn := range clients {
					go func(conn *client.Client) {
						conn.Send("Server is shutting down.")
						conn.Quit()
					}(conn)
				}
				shutdown = true
			}
			go cmd.Exec()
		case conn := <-add:
			clients[conn.Name()] = conn
		case rmConn := <-rm:
			log.Printf("%s disconnected", rmConn.Name())
			delete(clients, rmConn.Name())
			if shutdown && len(clients) == 0 {
				return
			}
		}
	}
}

func acceptConnections(srv net.Listener, rm, add chan *client.Client, msg chan command.Commander) {
	log.Printf("Server is now running on port: %v\n", srv.Addr())

	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Connection from: %v", conn.RemoteAddr())
		go client.Handle(conn, rm, add, msg)
	}
}
