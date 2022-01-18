# Solution to  Junior  Developer role -Http server


Writing an Http Server  in Golang.


The server serves like below:


![Server Demo](demo/tcpwalk.gif)


**Background**

     Design a basic HTTP web-server application which can listen on a configurable TCP port and serve both static HTML and 
     dynamically generated HTML by means of a chosen programming language, such as in the way Apache uses PHP. 
     It is acceptable for this server application to support only a restricted subset of HTTP, such as GET or POST requests, 
     and the only headers it must support are Content-Type and Content-Length.

**Prerequisite**

You need to have the following to run the server in your machine:
- A basic Understanding of [Golang](https://go.dev/) and how to set up your [Go workspace.](https://go.dev/doc/gopath_code)
- Go installed in your machine


**Environment Setup**

To ensure you have all the dependencies needed to run the application,run the following command
from the root folder of the project

        go mod tidy

 ## Building an Http Web server on a configurable tcp port  
 The server is basically started by running the command `go run main.go port`

 > Port here is a valid port number where the tcp server will listen to connections made.

 Once the client makes the connection to the right port,the http server starts serving 
 the required static html and dynamically generated html.

 A  [mux router](https://github.com/gorilla/mux) is encoded by the client then decoded by the connection handler then used to handle different requests.  
 
        > Client
        func Client(port string) {

	        c, err := net.Dial("tcp", port)

	        if err != nil {
		      fmt.Println(err)
		      fmt.Println("Here now")

	     }

	     //Inialize an instance of the Mux Router and encode it as A pointer
	      r := mux.NewRouter()
	     gob.NewEncoder(c).Encode(*r)
	     c.Close()

        }

        > Handler
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
## Demonstrating how the Server handles Requests(Post & Get)
To show how the server handles Get and Post requests,i have added a simple html form.One can add the name and email then submit.
The name is then added to the members' list as shown below:

![PostReq Demo](demo/posteget.gif)
## Securing your Http Server with a self signed Certificate
Use the openssl command below to create a private key (go-server.key) and a self-signed certificate (go-server.crt) valid for 365 days with a key size of 2,048 bits.
Please note that this should be run from the root directory of the project.

       openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout go-server.key -out go-server.crt


Depending on your organization name and current location, respond to each question appropriately. Make sure you've provided the correct public IP address or domain name when prompted to enter.

If everything went well,you should now see the `go-server.crt` and `go-server.key` files on the file list in your project.

Now visiting the server without a secure connection,you should see the message on your browser as below:

![Http Exception](demo/httpserr.PNG)