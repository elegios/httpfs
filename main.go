package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := "9000"
	var dir string
	switch len(os.Args) {
	case 3:
		port = os.Args[1]
		dir = os.Args[2]
	case 2:
		dir = os.Args[1]
	default:
		log.Fatal("Incorrect number of arguments. Need \"[port] directory\"")
	}

	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Printf("Hosting %s on port %s.", dir, port)
	http.ListenAndServe(":"+port, nil)
}
