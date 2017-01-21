package main

import (
    "fmt"
    "net/http"
)

func main()  {
    http.HandleFunc("/",handler1)
    http.HandleFunc("/route2",handler2)
    http.ListenAndServe(":8080",nil)
}

func handler1(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w,"Hi Browser!.. I am Golang Server")
}

func handler2(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "I am the routed url from Home Page")
}
