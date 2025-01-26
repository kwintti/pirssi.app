package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
    fileServer := http.FileServer(http.Dir("./static"))
    fmt.Println("Serving on: 8080")
    http.Handle("/", fileServer)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
