go build
./GoTunnel --client &
./GoTunnel --server --port=1234 --remote-port=9000 &
echo "Starting Netcat on port 9000. Open another console and type in nc 127.0.0.1 1212"
nc -l 9000 
rm GoTunnel