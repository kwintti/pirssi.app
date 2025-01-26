package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
    fileServer := http.FileServer(http.Dir("./static"))
    fmt.Println("Serving on: 80")
    http.Handle("/", fileServer)
    log.Fatal(http.ListenAndServe(":80", nil))
}


