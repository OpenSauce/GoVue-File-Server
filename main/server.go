package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/avaliablespace", avaliablespaceHandler)

	fs := http.FileServer(http.Dir("../frontend/dist/"))
	http.Handle("/", fs)


	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

// type thumbnailRequest struct {
// }

func avaliablespaceHandler(w http.ResponseWriter, r *http.Request) {
	// var decoded thumbnailRequest

	// // Try to decode the request into the thumbnailRequest struct.
	// err := json.NewDecoder(r.Body).Decode(&decoded)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	fmt.Printf("Recieved the message %f ", printUsage("/"))
	fmt.Fprintf(w, `{ "avaliablespace": "%f" }`, printUsage("/"))
}