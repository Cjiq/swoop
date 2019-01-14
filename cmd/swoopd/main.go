package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

type context struct {
	host string
	port int
}

func (c context) hostname() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

var con = &context{}

func init() {
	flag.StringVar(&con.host, "host", "localhost", "Set hostname for server")
	flag.StringVar(&con.host, "h", "localhost", "Set hostname for server")
	flag.IntVar(&con.port, "port", 5625, "Set port for server")
	flag.IntVar(&con.port, "p", 5625, "Set port for server")
	flag.Parse()
}

func main() {
	// Listen for incoming connections
	list, err := net.Listen("tcp", con.hostname())
	if err != nil {
		log.Fatal(err)
	}
	defer list.Close()

	log.Printf("Swoop Server Deamon Started: %s\n", con.hostname())
	for {
		conn, err := list.Accept()
		if err != nil {
			log.Fatalf("Failed to accept connection: %s\n", err.Error())
		}
		log.Printf("New connection from: %s\n", conn.RemoteAddr().String())
		go serveConnection(conn)
	}
}

func serveConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Welcome to Swoop\n"))
}
