package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/butlermatt/gmud/command"
	"github.com/butlermatt/gmud/lib"
)

// Client holds the connection and channel
type Client struct {
	// Conn is the raw network connection.
	Conn net.Conn
	// Ch channel is the text to send to the client.
	Ch chan string
	// Quit channel will force the client to disconnect.
	quit chan bool
	// Name is the username of the player.
	name string
	// Long description of the player.
	long string

	// rm channel is the channel to remove the client from the server.
	rm chan<- *Client
	// hasQuit flag indicates if the quit has been sent or not.
	hasQuit bool
	// room the user is currently located in.
	room *lib.Room
}

// Prompt sends the user a prompt (default is > otherwise specified prompt)
func (c *Client) Prompt(p string) {
	if p == "" {
		return
	}
	fmt.Fprint(c.Conn, "\r"+p)
}

// toClient manages incoming stream and redirects to the client socket appropriately.
func (c *Client) toClient() {
	defer c.Conn.Close()
	for {
		select {
		case msg := <-c.Ch:
			fmt.Fprintln(c.Conn, "\r"+msg)
			c.Prompt("> ")
		case <-c.quit:
			c.rm <- c
			return
		}
	}
}

// Write sends a string to the client in a non-blocking way.
func (c *Client) Write(str string) {
	go func() { c.Ch <- str }()
}

// Send sends a string to the client in a potentially blocking way.
func (c *Client) Send(str string) {
	c.Ch <- str
}

// Quit implements the Quiter command interface.
func (c *Client) Quit() {
	c.hasQuit = true
	c.quit <- true
}

// Name returns the name of the client. Fulfills the Objecter interface.
func (c *Client) Name() string {
	return c.name
}

// Description returns the description of the player.
func (c *Client) Description() string {
	return c.long
}

// Room returns a pointer to the current room occupied by the user.
func (c *Client) Room() lib.Holder {
	return c.room
}

// Handle creates a new Client from the connection and channel to server
func Handle(conn net.Conn, rm, add chan<- *Client, msg chan<- command.Commander) {
	fmt.Fprintln(conn, "Welcome to GMud!")

	client := &Client{
		Conn: conn,
		Ch:   make(chan string),
		quit: make(chan bool),
		rm:   rm,
	}

	go client.toClient()
	defer close(client.Ch)

	rd := bufio.NewReader(conn)

	var name string

	// TODO: Login Stuff
	for name == "" {
		client.Prompt("What is your name? ")
		nm, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("Error reading line: %v", err)
			break
		}
		name = strings.TrimSpace(string(nm))
	}

	if name == "" {
		log.Println("Error getting name")
		return
	}
	client.name = name
	add <- client

	lib.DefaultRoom.Add(client)
	client.room = lib.DefaultRoom

	client.Prompt("> ")
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		txt := scan.Text()
		wrds := strings.Fields(txt)

		if len(wrds) == 0 {
			client.Prompt("> ")
			continue
		}
		cmd, err := command.GetCommand(client, wrds)
		if err != nil {
			client.Write(err.Error())
			continue
		}
		msg <- cmd
	}
	if !client.hasQuit {
		log.Printf("Client exited without quit: %v\n", client.Conn.RemoteAddr())
		client.quit <- true
	}

	if err := scan.Err(); err != nil {
		log.Println(err)
	}
}
