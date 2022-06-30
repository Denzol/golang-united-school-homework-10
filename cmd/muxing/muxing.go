package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", handleFirst).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handleData).Methods(http.MethodPost)
	router.HandleFunc("/headers", handleHeaders).Methods(http.MethodPost)
	http.Handle("/", router)
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func handleFirst(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintf(writer, "Hello, %s!", vars["PARAM"])
}

func handleBad(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusInternalServerError)
}

func handleData(writer http.ResponseWriter, request *http.Request) {
	data, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(writer, "I got message:\n%s", string(data))
}

func handleHeaders(writer http.ResponseWriter, request *http.Request) {
	val1, err := strconv.Atoi(request.Header["a"][0])
	if err != nil {
		log.Fatal(err)
	}
	val2, err := strconv.Atoi(request.Header["b"][0])
	if err != nil {
		log.Fatal(err)
	}
	value := val1 + val2
	b := strconv.Itoa(value)
	writer.Header().Add("a+b", b)
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "8081")
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
