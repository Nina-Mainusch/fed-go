package client

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {

	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	clients := []string{"client1", "client2"} //
	data := []string{"data", "data2"}
	var parameters [][]byte

	// Collect all the parameters
	for i := range clients {
		fmt.Println(i, clients[i], data[i])
		coordinates := GetCoordinates("../../data/" + data[i] + ".csv")
		// Do linear regression
		aHat, bHat := GetCoefficients(coordinates)
		parameters = append(parameters, float64ToByte(aHat), float64ToByte(bHat))
		fmt.Println("Original:", aHat)
		fmt.Println("Original:", bHat)
		// Send parameters to server
		_, err = connection.Write(float64ToByte(aHat))
		_, err = connection.Write(float64ToByte(bHat))
	}

	// Receive answer from server
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Client Received: ", Float64frombytes(buffer[:mLen]))
	defer connection.Close()
}
