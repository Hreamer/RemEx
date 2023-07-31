package main

import (
	"flag"
	"fmt"
)

/*
	Tasks:
		Task 1: Use flags library to make Help flag [X]
		Task 2: Get File via args [X]
		Task 3: Open the File []
		Task 4: Send the file []
		Task 5: Wait for the return []
		Task 6: Grab the return []
		Task 7: Print the return []

*/

func main() {
	//This will Parse the flags and also print help message for all flags if -help is called
	var path = flag.String("PATH", "", "The path of the file you wish to send")
	flag.Parse()

	fmt.Printf("-PATH was set to: %s\n", *path)

}
