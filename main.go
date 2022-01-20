package main

import (
	"log"
	"os"
	"net/http"
	"fmt"
)

func main() {
	port := "8080"
	path := "."
	args := os.Args[1:]
	path_found := false
	for i := 0; i < len(args); i+=1 {
		if args[i] == "-p" {
			if i < len(args)-1 {
				i+=1
				port = args[i]
			} else {
				fmt.Printf("Expected the port after \"-p\"\n")
				return
			}
		} else if !path_found {
			path_found = true
			path = args[i]
		} else {
			fmt.Printf("Extra argument \"%s\"\n", args[i])
			return
		}
	}

	fmt.Printf("Serving directory \"%s\" on port %s\n", path, port)
	// Simple static webserver:
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, http.FileServer(http.Dir(path))))
}


