package main

import (
	"flag"
	"log"
	"net"
)

var (
	ip_address        = flag.String("ip-address", "127.0.0.1", "Please attach a valid IP address")
	port              = flag.String("port", "1212", "The port you wish to use")
	remote_ip_address = flag.String("remote-ip-address", "127.0.0.1", "Please attach a valid IP address")
	remote_port       = flag.String("remote-port", "1234", "The port you wish to use")
	public_key        = flag.String("public-key", "public.pem", "Please enter the path of your public key")
	private_key       = flag.String("private-key", "private.pem", "Please enter the path of your private key")
	create_cert       = flag.Bool("create-cert", false, "Create Public and Private PEM")
	server            = flag.Bool("server", false, "You are accepting TLS connections from other hosts")
	client            = flag.Bool("client", false, "You are tunneling connections to a server")
)

func main() {
	flag.Parse()
	if *create_cert {
		CreateEncryptionKeys()
		return
	}

	addr := net.ParseIP(*ip_address)
	if addr == nil {
		log.Fatalln("Unable to parse IP. IP Provided was", *ip_address)
	}
	service := *ip_address + ":" + *port

	raddr := net.ParseIP(*remote_ip_address)
	if raddr == nil {
		log.Fatalln("Unable to parse IP. IP Provided was", *remote_ip_address)
	}
	RemoteIPandPort := *remote_ip_address + ":" + *remote_port

	var listener net.Listener
	var secure bool

	if *server {
		listener = ServeTLSConnections(*public_key, *private_key, service)
		secure = false // We will be connecting to our service using TCP
	} else { // if we are not a server we must be a client
		listener = ServeTCPConnections(service)
		secure = true // We need to connect using TLS over the wire
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		log.Println("Connection Accepted from", conn.RemoteAddr().String())
		go handleClient(conn, RemoteIPandPort, secure)
	}
}
