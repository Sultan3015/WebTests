package main

import (
	"fmt"
	"log"
	"net/http"
)

//Responswriter value assembles the HTTP server's response; by writing to it , we send data to the HTTP client
//http.Request represents is the data structure that represents the clients HTTP request
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, I love this shit %s!", r.URL.Path[1:]) //URL.Path is the path component of the URL request
}

//Main function calls the http.handleFunc, ehich tells the http package to handle all the requests to the web root with handler
//it then calls https.listenAndServe, specifying tha it should listen on port 8080
func main() {
	http.HandleFunc("/", handler) //web root with handler
	log.Fatal(http.ListenAndServe(":8080", nil))
	//ListenAndServe always returns an error but wrapping it using the log.Fatal logs that error
}
