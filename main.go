package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var store = flag.String("store", "", "Location of the files")
var action = flag.String("action", "", "Action over the file")
var file = flag.String("file", "", "New file name")

func readFile() {
	data, err := ioutil.ReadFile(*store + "/" + *file)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nData: %s", data)
}

func handleRequests() {
	flag.Parse()

	if *action == "read" {
		readFile()
	}

	http.Handle("/", http.FileServer(http.Dir(*store)))

	port := ":8000"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}

func main() {
	handleRequests()
}
