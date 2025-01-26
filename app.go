package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
    //fileServer := http.FileServer(http.Dir("./static"))
    fmt.Println("Serving on: 80")
    //http.Handle("/", fileServer)
    http.HandleFunc("/", page)
    log.Fatal(http.ListenAndServe(":80", nil))
}

func page(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "<html><body><h1>Hello!</h1></body></html>")
}
