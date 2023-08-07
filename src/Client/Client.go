package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
	Tasks:
		Task 1: Use flags library to make Help flag [X]
		Task 2: Get File via args [X]
		Task 3: Open the File [X]
		Task 4: Send the file []
		Task 5: Wait for the return []
		Task 6: Grab the return []
		Task 7: Print the return []

*/

func main() {
	//Parse the flags and also print help message for all flags if -help is called
	var path = flag.String("p", "", "The path of the file you wish to send")
	var address = flag.String("a", "", "The address to send the file")
	var language = flag.String("l", "", "The language of the file. Check README for supported languages")

	flag.Parse()

	//Flag Test Code
	//fmt.Printf("-PATH was set to: %s\n", *path)
	//fmt.Printf("-Address was set to %s\n", *address)
	//fmt.Printf("-Language was set to: %s\n", *language)

	//open the file
	file, err := os.Open(*path)
	if err != nil {
		fmt.Printf("RemEx could not open file in path %s\n", *path)
		os.Exit(0)
	}

	//send the file
	//Note: os.Open returns a file type and it implements the io.Reader interface
	resp, postErr := http.Post(*address+*language, "text/plain", file)
	if postErr != nil {
		fmt.Println("RemEx failed to send the POST request")
		os.Exit(0)
	}

	//Get the Body as a string slice so we can convert it to string to print
	body, ioerr := io.ReadAll(resp.Body)
	if ioerr != nil {
		fmt.Println("Could not read the body of the response")
		os.Exit(0)
	}
	resp.Body.Close()

	//Print reponse
	fmt.Printf("Server Response:\n%s\n", string(body))

	file.Close()
}
