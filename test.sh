./GoTunnel --client &
./GoTunnel --server --port=1234 --remote-port=9000 &
echo "Starting Netcat on port 9000"
nc -l 9000 

