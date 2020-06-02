package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jsfong/sample-http-go-server/pkg/echoer"
)

func echoGenericRequest(r *http.Request) {
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	body := string(bs)

	header := r.Header

	fmt.Println("-----Request START-------")
	fmt.Println("Request: ", r)
	fmt.Println("-----Request END-------")
	fmt.Println()
	fmt.Println("-----Header START-------")
	fmt.Println("Request Header: ", header)
	fmt.Println("-----Header END-------")
	fmt.Println()
	fmt.Println("-----Body START-------")
	fmt.Println("Body: ", body)
	fmt.Println("-----Body END-------")
	fmt.Println()
	fmt.Println("----------------------------------------")

}

func echoAndServeXMLFile(w http.ResponseWriter, r *http.Request) {

	echoer.EchoMultipart3(r)

	filename := "pkg/response/sample.xml"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to parse file %s", filename)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(b))
}

func echoAndServeJSONFile(w http.ResponseWriter, r *http.Request) {

	echoer.EchoMultipart3(r)

	filename := "pkg/response/sample.json"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to parse file %s", filename)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(b))

}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/check_liveness", echoAndServeJSONFile)

	// EXAMPLE
	// router.HandleFunc("/check_liveness", echoAndServeXMLFile)
	// // router.HandleFunc("/event", createEvent).Methods("POST")
	// // router.HandleFunc("/events", getAllEvents).Methods("GET")
	// // router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	// // router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	// // router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
