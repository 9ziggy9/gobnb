package main

import (
  "fmt"
  "log"
  "github.com/gorilla/mux"
  "net/http"
  "os"
  "github.com/joho/godotenv"
)

func init() {
  err := godotenv.Load()
  if err != nil {
    defer func() {
      log.Fatal("Error loading .env file")
    }()
    panic(err)
  }
}

func main() {
  PORT := os.Getenv("PORT")
  // DB := os.Getenv("DB")

  // Initialize gorilla router
  r := mux.NewRouter()

  // ROUTE WITH ANON HANDLER
  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World!")
  }).Methods("GET")

  fmt.Println(fmt.Sprintf("Starting server on port %s ...", PORT))
  err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
  if err != nil {
    defer func() {
      log.Println("Hit deferred function, panicked out")
      log.Fatal(err)
    }()
    panic(err)
  }
}
