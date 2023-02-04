package server

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		// The server needs to wait for clients to send him updates
		// as soon as he received two updates, it should compute the average and return it to the senders
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(connection)
	}
}
func processClient(connection net.Conn) {
	offset := 8
	buffer := make([]byte, 1024)
	var data [][]byte
	// We need to wait till all the clients send their updates
	for {
		mLen, err := connection.Read(buffer)
		fmt.Println("Len of buffer:", mLen)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		for i := 0; i < mLen; i = i + offset {
			fmt.Println("Server Received: ", Float64frombytes(buffer[i:i+offset]))
			_, err = connection.Write(buffer[i : i+offset])
			data = append(data, buffer[i:i+offset])
		}
		if len(data) > 3 {
			// TODO now we gotta compute the averages and return those instead!
			// Ideally what I want is different clients sending their updates in their own TCP connection.
			// Right now all client updates go through the same pipeline.
			break
		}
	}
	connection.Close()
}
