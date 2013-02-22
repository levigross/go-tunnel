package main

import (
	"crypto/rand"
	"crypto/tls"
	"log"
	"net"
	"time"
)

func ServeTCPConnections(service string) net.Listener {
	listener, err := net.Listen("tcp", service)
	checkError(err)
	log.Println("Listening for TCP connections on", service)
	return listener
}

func ServeTLSConnections(public_key, private_key, service string) net.Listener {
	cert, err := tls.LoadX509KeyPair(public_key, private_key)
	checkError(err)

	config := tls.Config{Certificates: []tls.Certificate{cert}}

	now := time.Now()
	config.Time = func() time.Time { return now }
	config.Rand = rand.Reader

	listener, err := tls.Listen("tcp", service, &config)
	checkError(err)
	log.Println("Listening for TLS connections on", service)
	return listener
}

func handleClient(conn net.Conn, remoteServerIPandPort string, secure bool) {
	defer conn.Close()

	var buf [512]byte
	var server_conn net.Conn
	if secure {
		server_conn, err := tls.Dial("tcp", remoteServerIPandPort, &tls.Config{InsecureSkipVerify: false})
		checkError(err)
		defer server_conn.Close()
	} else {
		server_conn, err := net.Dial("tcp", remoteServerIPandPort)
		checkError(err)
		defer server_conn.Close()
	}

	for {
		log.Println("Trying to read")
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Println(err)
			continue
		}
		_, err2 := server_conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}
