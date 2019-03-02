// Declare current package to be main
package main

// This imports several libraries in bulk using 'import ( ... )' ?
import (
  "net/http" // For handling HTTP requests
  "strings" // For string manipulation
)

// Recieve HTTP requests on route /
// Sends back any portion of the URL after localhost
func sayHello(writer http.ResponseWriter, request *http.Request) {

  // 'writer' seems to be roughly anlogous to 'res'
  // 'request' seems to be rougly analogous to 'req'
  // (Or I could just name it 'req' if I wanted to)

  // Assigns the path of the request URL to a variable using ':='
  requestPath := request.URL.Path
  name := strings.TrimPrefix(requestPath, "/")
  message := "Hello " + name

  // Sends 'message' back to the request sender
  writer.Write([]byte(message))
}

// Main function for the main file 'main.go'
func main() {

  // Handle HTTP requests on /
  // Unlike node.js/express this seems to handle all types of HTTP requests
  http.HandleFunc("/", sayHello)

  // If there is an error generated log it
  // Panic() seems to end the program
  if err := http.ListenAndServe(":8000", nil); err != nil {
    panic(err)
  }
}