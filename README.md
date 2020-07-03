## Small program to interact with the file system 

Files can be read and created via the command line. 
`http.FileServer` can be used to display the contents of your store and files in the browser.

To start the server and set up the store where the files that you'll work with are located: 
 `go run main.go --store=./tmp`
 
Visit localhost to see the contents of your store
 `http://localhost:8000/`

See the contents of a file in the terminal (where file2.txt is the name of the file):
`go run main.go --store=./tmp --action=read --file=file1.txt`

Create a new file and write to it:
`go run main.go --store=./tmp --action=create --file=file3.txt --content=some_content`

Update the content of a file (this will overwrite any content that is already present in the file)
`go run main.go --store=./tmp --action=update --file=file3.txt --content=some_other_content`