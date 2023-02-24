package main

import (
  "fmt"
  "log"
  "github.com/gorilla/mux"
  "net/http"
)

func main() {
  r := mux.NewRouter()

  // ROUTE WITH ANON HANDLER
  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World!")
  }).Methods("GET")

  fmt.Println("Starting server on :8000 ...")
  err := http.ListenAndServe(":8000", r)
  if err != nil {
    defer func() {
      log.Println("Hit deferred function, panicked out")
      log.Fatal(err)
    }()
    panic(err)
  }
}
