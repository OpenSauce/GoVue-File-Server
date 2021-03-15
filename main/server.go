package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/avaliablespace", avaliablespaceHandler)
	http.HandleFunc("/api/executecommand", executeCommandHandler)

	fs := http.FileServer(http.Dir("../frontend/dist/"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 8080")
	log.Panic(
		http.ListenAndServe(":8080", nil),
	)
}

func avaliablespaceHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	stats := GetHDDStats("/")

	fmt.Printf("Recieved the message %f \n", GetHDDStats("/").Percentage)
	fmt.Fprintf(w, `{ "avaliablespace": "%f",
	"totalSpace": "%s",
	"freeSpace": "%s",
	"pcName": "%s" }`, stats.Percentage, stats.TotalSpace, stats.FreeSpace, GetHostname())
}

type commandRequest struct {
	Command string `json:"command"`
}

func executeCommandHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var request commandRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Recieved the message %s \n", request.Command)
	output, err := ExecuteCommand(request.Command)
	formattedErr := ""
	if err != nil {
		formattedErr = err.Error()
	}

	fmt.Printf("Got: %s : %s \n", output, formattedErr)
	fmt.Fprintf(w, `{ "output": "%s",
	"error": "%s" }`, output, formattedErr)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
