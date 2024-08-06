package main

import (
 "fmt"
 "net/http"
 "time"

 "github.com/lemrex/go1/handler"
)

func main() {
 http.HandleFunc("/", handler.Welcome)
 http.HandleFunc("/time", handler.Time)
 http.HandleFunc("/date", handler.Date)  
 http.HandleFunc("/day", handler.Day)     
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