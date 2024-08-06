package handler

import (
 "fmt"
 "net/http"
 "os"
 "time"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome:", os.Getenv("MYNAME"))
   }

func Hello(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintln(w, "Hello, World! My name is:", os.Getenv("MYNAME"))
}

func Goodbye(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintln(w, "Goodbye, World! My name is:", os.Getenv("MYNAME"))
}


func Time(w http.ResponseWriter, r *http.Request){
	time:=time.Now()
	fmt.Fprintln(w, "The time is:", time)
		
}

func Day(w http.ResponseWriter, r *http.Request){
	day:=time.Now().Local().Weekday()
	fmt.Fprintln(w, "Today is:", day)

}

func Date(w http.ResponseWriter, r *http.Request){
	date:=time.Now().Format("2006-01-02")
	fmt.Fprintln(w, "The date is:", date)

}