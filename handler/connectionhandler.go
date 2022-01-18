package handler

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

/*

   Once the client dials into the connection with the right port,the http server starts serving content
   based on the matched url.[using mux router]

*/

var members []Pesapal

func Client(port string) {

	c, err := net.Dial("tcp", port)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Here now")

	}
	fmt.Println("Connection Successfull")
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
		fmt.Println(err)
	} else {

		//Static file(s) configuration
		staticFileDirectory := http.Dir("./assets/")
		staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

		//Get requests
		r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
		r.HandleFunc("/bird", GetConfirmation).Methods("GET")
		r.HandleFunc("/members", GetMembersHandler).Methods("GET")

		//Post requests
		r.HandleFunc("/bird", CreateMember).Methods("POST")

		//Server Configurations
		srv := &http.Server{
			Handler:      r,
			Addr:         "127.0.0.1:8000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		//Securing the server with a self-Signed Certificate
		srv.ListenAndServeTLS("go-server.crt", "go-server.key")
	}
	c.Close()

}

func GetConfirmation(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content Type", "text/html")
	// The template name "template" does not matter here
	templates := template.New("template")
	templates, _ = templates.ParseFiles("assets/index.html")
	// "doc" is the constant that holds the HTML content
	//doc := "./assets"
	//templates.New("doc").Parse()

	pesapal := Pesapal{
		Name:  "Android Engineers",
		Email: "Tech Company",
	}
	//templates.Lookup("doc").Execute(w, pesapal)
	templates.Execute(w, pesapal)
}

func CreateMember(w http.ResponseWriter, r *http.Request) {

	// Create a new instance of Pasapal(structure)
	member := Pesapal{}

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	member.Name = r.Form.Get("name")
	member.Email = r.Form.Get("email")

	members = append(members, member)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func GetMembersHandler(w http.ResponseWriter, r *http.Request) {

	//Convert the "members" variable to json
	membersList, err := json.Marshal(members)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of members to the response
	w.Write(membersList)
}

type Pesapal struct {
	Name  string
	Email string
}
