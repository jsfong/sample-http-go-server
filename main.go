package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jsfong/sample-http-go-server/echoer"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func readAndPrintBody(w http.ResponseWriter, r *http.Request) {
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	body := string(bs)

	header := r.Header

	fmt.Println("Request Header: ", header)
	fmt.Println("Request Body: ", body)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", echoer.EchoMultipart3)
	// // router.HandleFunc("/event", createEvent).Methods("POST")
	// // router.HandleFunc("/events", getAllEvents).Methods("GET")
	// // router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	// // router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	// // router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
