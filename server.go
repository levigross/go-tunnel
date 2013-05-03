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
	checkError(err, true)
	log.Println("Listening for TCP connections on", service)
	return listener
}

func ServeTLSConnections(public_key, private_key, service string) net.Listener {
	cert, err := tls.LoadX509KeyPair(public_key, private_key)
	checkError(err, true)

	config := tls.Config{Certificates: []tls.Certificate{cert}}

	now := time.Now()
	config.Time = func() time.Time { return now }
	config.Rand = rand.Reader

	listener, err := tls.Listen("tcp", service, &config)
	checkError(err, true)
	log.Println("Listening for TLS connections on", service)
	return listener
}

func GetConnection(remoteServerIPandPort string, secure bool) net.Conn {

	if !secure {
		server_conn, err := net.Dial("tcp", remoteServerIPandPort)
		checkError(err, true)
		return server_conn
	}

	cert, err := tls.LoadX509KeyPair("public.pem", "private.pem")
	checkError(err, true)

	config := tls.Config{Certificates: []tls.Certificate{cert}, Time: func() time.Time { return time.Now() },
		Rand: rand.Reader, InsecureSkipVerify: true}

	server_conn, err := tls.Dial("tcp", remoteServerIPandPort, &config)
	checkError(err, true)
	return server_conn

}

func handleClient(conn net.Conn, remoteServerIPandPort string, secure bool) {
	defer conn.Close()

	server_conn := GetConnection(remoteServerIPandPort, secure)
	defer server_conn.Close()

	var in_buf [512]byte
	var out_buf [512]byte

	for {
		n, err := conn.Read(in_buf[0:])
		if err != nil {
			log.Println(err)
			return
		}
		_, err2 := server_conn.Write(in_buf[0:n])
		if err2 != nil {
			return
		}
		o, err3 := server_conn.Read(out_buf[0:])
		if err3 != nil {
			log.Println(err3)
			return
		}
		_, err4 := conn.Write(out_buf[0:o])
		if err4 != nil {
			return
		}

	}
}
