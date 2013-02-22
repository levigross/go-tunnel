Go Tunnel
========


This is a program like STunnel but written in Go. Most of the code is derived from [Network programming with Go](http://jan.newmarch.name/go/) and I take no credit for his outstanding work.

Usage
====
Usage of ./go-tunnel:
  -client=false: You are tunneling connections to a server
  -create-cert=false: Create Public and Private PEM
  -ip-address="127.0.0.1": Please attach a valid IP address
  -port="1212": The port you wish to use
  -private-key="private.pem": Please enter the path of your private key
  -public-key="public.pem": Please enter the path of your public key
  -remote-ip-address="127.0.0.1": Please attach a valid IP address
  -remote-port="1234": The port you wish to use
  -server=false: You are accepting TLS connections from other hosts