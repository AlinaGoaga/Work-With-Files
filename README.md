## Small program to read/create files via the command line 
### http.Server can be used to display the contents of your store / files in the browser

To start server: 
 `go run main.go --store=./tmp`
 
Visit localhost to see the contents of your store
 `http://localhost:8000/`

To see the contents of a file in the terminal (where file2.txt is the name of the file)
`go run main.go --action=read --file=file2.txt`
