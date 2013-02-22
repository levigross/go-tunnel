#!/bin/sh
go build
BIN_NAME="$(basename `pwd` )"
./$BIN_NAME --client &
./$BIN_NAME --server --port=1234 --remote-port=9000 &
echo "Starting Netcat on port 9000. Open another console and type in nc 127.0.0.1 1212"
nc -l 9000 
rm $BIN_NAME