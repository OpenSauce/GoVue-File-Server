package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/OpenSauce/GoVue-File-Server/pkg/filesystem"
)

type endpoint struct {
	EndpointURL string
	Handler     func(http.ResponseWriter, *http.Request)
}

var endpointList = []endpoint{
	{EndpointURL: "/api/avaliablespace", Handler: avaliablespaceHandler},
	{EndpointURL: "/api/executecommand", Handler: executeCommandHandler},
	{EndpointURL: "/api/upload", Handler: uploadHandler},
	{EndpointURL: "/api/getfiles", Handler: getFileHandler},
}

//Start the HTTP Server.
func Start(doneCh chan struct{}) {
	for _, endpoint := range endpointList {
		http.HandleFunc(endpoint.EndpointURL, endpoint.Handler)
	}

	fs := http.FileServer(http.Dir("../frontend/dist/"))
	http.Handle("/", fs)

	log.Println("Server listening on port 8080")
	log.Panic(
		http.ListenAndServe(":8080", nil),
	)
}

//Endpoint for returning the avaliable space.
func avaliablespaceHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	stats := filesystem.GetHDDStats("/")

	log.Printf("Recieved the message %f \n", filesystem.GetHDDStats("/").Percentage)
	fmt.Fprintf(w, `{ "avaliablespace": "%f",
	"totalSpace": "%s",
	"freeSpace": "%s",
	"pcName": "%s" }`, stats.Percentage, stats.TotalSpace, stats.FreeSpace, filesystem.GetHostname())
}

//Endpoint for returning the list of files.
func getFileHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	listOfFiles := filesystem.GetListOfFiles("/tmp/")
	filesJson, _ := json.Marshal(listOfFiles)

	log.Printf(`{ "files": "%s" }`, string(filesJson))
	fmt.Fprintf(w, `{ "files": %s }`, string(filesJson))
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

	log.Printf("Recieved the message %s \n", request.Command)
	output, err := filesystem.ExecuteCommand(request.Command)
	formattedErr := ""
	if err != nil {
		formattedErr = err.Error()
	}

	log.Printf("Got: %s : %s \n", output, formattedErr)
	fmt.Fprintf(w, `{ "output": "%s",
	"error": "%s" }`, output, formattedErr)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//Handler for handling the uploading of multiple files
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	err := r.ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	formdata := r.MultipartForm // ok, no problem so far, read the Form data

	//get the *fileheaders
	files := formdata.File["multiplefiles"] // grab the filenames

	for i := range files { // loop through the files one by one
		file, err := files[i].Open()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		defer file.Close()

		out, err := os.Create("/tmp/" + files[i].Filename)
		if err != nil {
			log.Printf("Unable to create the file for writing. Check your write access privilege")
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file) // file not files[i] !

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		log.Printf("Files uploaded successfully : ")
		log.Printf(files[i].Filename + "\n")

	}
}
