package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	//Set up handlers
	http.HandleFunc("/go", goHandler)

	//Start listening for requests
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func goHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received /Go Request")

	if r.URL.Path != "/go" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	//read file contents
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could Not Read Body", http.StatusNotFound)
		return
	}

	//Create file to store the contents of the POST request
	//NOTE: If test.go exists it will delete the contents of this file so no worries
	file, createErr := os.Create("test.go")
	if createErr != nil {
		http.Error(w, "Could Not Open File", http.StatusNotFound)
		return
	}

	//Now write the contents to the file
	file.Write(body)
	file.Close()

	cmd := exec.Command("go", "run", file.Name())
	stdout, err := cmd.Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, string(stdout))
}
