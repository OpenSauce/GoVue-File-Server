package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	endpointList = []struct {
		EndpointURL string
		Handler func(http.ResponseWriter, *http.Request)
	}{
		{EndpointURL: "/api/avaliablespace", Handler: avaliablespaceHandler},
		{EndpointURL: "/api/executecommand", Handler: executeCommandHandler},
		{EndpointURL: "/api/upload", Handler: uploadHandler},
		{EndpointURL: "/api/getfiles", Handler: getFileHandler},
	}
)

//Main execution function of the server.
func main() {
	CreateDirectory("/gv-repo/")

	for _, endpoint := range endpointList {
		http.HandleFunc(endpoint.EndpointURL, endpoint.Handler)
	}

	fs := http.FileServer(http.Dir("../frontend/dist/"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 8080")
	log.Panic(
		http.ListenAndServe(":8080", nil),
	)
}

//Endpoint for returning the avaliable space.
func avaliablespaceHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	stats := GetHDDStats("/")

	fmt.Printf("Recieved the message %f \n", GetHDDStats("/").Percentage)
	fmt.Fprintf(w, `{ "avaliablespace": "%f",
	"totalSpace": "%s",
	"freeSpace": "%s",
	"pcName": "%s" }`, stats.Percentage, stats.TotalSpace, stats.FreeSpace, GetHostname())
}

//Endpoint for returning the list of files.
func getFileHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	listOfFiles := GetListOfFiles("gv-repo/")
	filesJson, _ := json.Marshal(listOfFiles)

	fmt.Printf(`{ "files": "%s" }`,string(filesJson))
	fmt.Fprintf(w, `{ "files": "%s" }`,string(filesJson))
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

//Handler for handling the uploading of multiple files
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Printf("Tracer 1")

	err := r.ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	formdata := r.MultipartForm // ok, no problem so far, read the Form data

	//get the *fileheaders
	files := formdata.File["multiplefiles"] // grab the filenames

	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		fmt.Printf("Name: %s", files[i].Filename)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		defer file.Close()

		out, err := os.Create("gv-repo/" + files[i].Filename)

		if err != nil {
			fmt.Printf("Unable to create the file for writing. Check your write access privilege")
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file) // file not files[i] !

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		fmt.Printf("Files uploaded successfully : ")
		fmt.Printf(files[i].Filename+"\n")

	}
}
