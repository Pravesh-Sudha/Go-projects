package main

import (
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))

}

func main(){

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/health", healthCheckHandler)

	port := "8090"
	fmt.Printf("Starting server on port %s....", port)

	err := http.ListenAndServe(":"+ port, nil)

	if err != nil {
		fmt.Println("Error starting server:",err)
	}

}