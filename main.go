package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var store = flag.String("store", "", "Location of the files")
var action = flag.String("action", "", "Action over the file")
var file = flag.String("file", "", "New file name")
var content = flag.String("content", "", "Content to be added to the file")

func readFile() {
	data, err := ioutil.ReadFile(*store + "/" + *file)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nData: %s\n", data)
}

func createFile() {
	data := []byte(*content)
	err := ioutil.WriteFile(*store + "/" + *file, data, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func updateContent() {      
	file, err := os.OpenFile(
        *store + "/" + *file,
        os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
        0666,
    )
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    byteContent := []byte(*content)
    file.Write(byteContent)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("The file has been succesfully updated")
}

func handleRequests() {
	flag.Parse()

	switch *action {
		case "read":
			readFile()
		case "create":
			createFile()
		case "update":
			updateContent()
		default:
			fmt.Println("Please choose an action in the CLI")
    }

	http.Handle("/", http.FileServer(http.Dir(*store)))

	port := ":8000"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}

func main() {
	handleRequests()
}
