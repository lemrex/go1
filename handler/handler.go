package handler

import (
 "fmt"
 "net/http"
 "os"
)

func Hello(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintln(w, "Hello, World! My name is:", os.Getenv("MYNAME"))
}

func Goodbye(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintln(w, "Goodbye, World! My name is:", os.Getenv("MYNAME"))
}