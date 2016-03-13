package main

import (
  "io"
  "net/http"
)

func serveRequest(responseWriter http.ResponseWriter, request *http.Request) {
  io.WriteString(responseWriter, "Hello world!")
}

func main() {
  http.HandleFunc("/", serveRequest)
  http.ListenAndServe(":8000", nil)
}
