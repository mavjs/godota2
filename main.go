package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/teams", teams)
	mux.HandleFunc("/players", players)
	var port = os.Getenv("PORT")
	if port == "" {
		port = ":" + "8080"
	} else {
		port = ":" + port
	}
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	fmt.Println("Started web server at ", server.Addr)
	server.ListenAndServe()
}
