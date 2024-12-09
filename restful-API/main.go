package main 

import (
	"log"
	"net/http"
	"project/tools"
)

func main(){
	http.HandleFunc("/tools", tools.ListTools)
	http.HandleFunc("/tools/add", tools.AddTool)
	http.HandleFunc("/tools/details", tools.ToolDetails)

	port := "8090"
	log.Printf("Starting the server at port %v", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting the server:",err)
	}
}