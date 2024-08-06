package main

import (
 "fmt"
 "net/http"
 "time"

 "example/handler"
)

func main() {
 http.HandleFunc("/hello", handler.Hello)
 http.HandleFunc("/goodbye", handler.Goodbye)

 server := &http.Server{
  Addr:         ":8080",
  Handler:      http.DefaultServeMux,
  ReadTimeout:  10 * time.Second, // Set a reasonable read timeout
  WriteTimeout: 10 * time.Second, // Set a reasonable write timeout
 }

 fmt.Println("Server started at :8080")
 if err := server.ListenAndServe(); err != nil {
  panic(err) // Handle the error appropriately
 }
}