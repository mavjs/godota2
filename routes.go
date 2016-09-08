package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	routes := map[string]string{
		"teams":   fmt.Sprintf("%s%s", request.Host, "/teams"),
		"players": fmt.Sprintf("%s%s", request.Host, "/players"),
	}
	if err := json.NewEncoder(writer).Encode(routes); err != nil {
		panic(err)
	}
}

func players(writer http.ResponseWriter, request *http.Request) {
}

func teams(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Teams Route: %s/%s", request.Host, request.URL.Path[1:])
}
