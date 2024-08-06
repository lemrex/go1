package handler

import (
 "fmt"
 "net/http"
 "net/http/httptest"
 "os"
 "testing"
)

func TestHelloHandler(t *testing.T) {
 os.Setenv("MYNAME", "the hash slinging slasher")
 req, err := http.NewRequest("GET", "/hello", nil)
 if err != nil {
  t.Fatal(err)
 }

 recorder := httptest.NewRecorder()
 handler := http.HandlerFunc(Hello)

 handler.ServeHTTP(recorder, req)

 if status := recorder.Code; status != http.StatusOK {
  t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
 }

 expected := fmt.Sprintf("Hello, World! My name is: %s\n", os.Getenv("MYNAME"))
 if recorder.Body.String() != expected {
  t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), expected)
 }
}

func TestGoodbyeHandler(t *testing.T) {
 os.Setenv("MYNAME", "the hash slinging slasher")

 req, err := http.NewRequest("GET", "/goodbye", nil)
 if err != nil {
  t.Fatal(err)
 }

 recorder := httptest.NewRecorder()
 handler := http.HandlerFunc(Goodbye)

 handler.ServeHTTP(recorder, req)

 if status := recorder.Code; status != http.StatusOK {
  t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
 }

 expected := fmt.Sprintf("Goodbye, World! My name is: %s\n", os.Getenv("MYNAME"))
 if recorder.Body.String() != expected {
  t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), expected)
 }
}