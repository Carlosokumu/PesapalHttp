package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Carlosokumu/PesapalTcp.git/handler"
)

func main() {

	/*
		Ensures that the user runs the program by providing a
		port number

	*/
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]

	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Start One")
		return
	}
	defer l.Close()

	for {

		c, err := l.Accept()

		if err != nil {
			fmt.Println(err)
			return

		}

		//Start the HandleServerConnection and the Client methods as goroutines to allow concurrency
		go handler.HandleServerConnection(c)
		go handler.Client(PORT)

	}

}
