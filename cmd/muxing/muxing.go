package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", Name).Methods(http.MethodGet)
	router.HandleFunc("/bad", Bad).Methods(http.MethodGet)
	router.HandleFunc("/data", Data).Methods(http.MethodPost)
	router.HandleFunc("/headers", Headers).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func Name(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Tester!")
}

func Bad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func Data(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "I got message:\n%s", data)
}

func Headers(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		log.Fatal(err)
	}
	Sum := a + b
	MySum := strconv.Itoa(Sum)
	w.Header().Add("a+b", MySum)
}
