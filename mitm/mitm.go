package mitm

import (
	"fmt"
	"log"
	"net"
)





func TestingMain(SERVER_NUM string, SERVER_TYPE string, SERVER_HOST string, SERVER_PORT string) (*error) {

	log.Printf("Testing Server %s Starting | %s | %s:%s", SERVER_NUM, SERVER_TYPE, SERVER_HOST, SERVER_PORT)

	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
			log.Fatal("Error listening: ", err.Error())
	}

	defer server.Close()
	
	log.Printf("Testing Server %s Started | %s | %s:%s", SERVER_NUM, SERVER_TYPE, SERVER_HOST, SERVER_PORT)

	for {
			connection, err := server.Accept()
			if err != nil {
					log.Fatal("Error Accepting: ", err.Error())
			}
			log.Printf("Client has connected")
			go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		log.Fatal("Error Reading: ", err.Error())
	}
	fmt.Printf("Received | %s", string(buffer[:mLen]))
	_, err = connection.Write([]byte("Message Recieved: " + string(buffer[:mLen])))
	if err != nil {
		log.Fatal("Error Responding: ", err.Error())
	}
	connection.Close()
}