package handler

import (
	"encoding/gob"
	"log"
	"net"

	"github.com/gorilla/mux"
)

/*

   Once the client dials into the connection with the right port,the http server starts serving content
   based on the matched url.[using mux router]

*/
func Client(port string) {

	c, err := net.Dial("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	//Inialize an instance of the Mux Router and encode it as A pointer
	r := mux.NewRouter()
	gob.NewEncoder(c).Encode(*r)
	c.Close()

}

/*
   It will listen to connections made and handle them.
   In this case,it will decode the router passed which is then  used to set up all
   the routes
*/

func HandleServerConnection(c net.Conn) {

	var r *mux.Router
	err := gob.NewDecoder(c).Decode(&r)

	if err != nil {
		log.Fatal(err)
	}
}
