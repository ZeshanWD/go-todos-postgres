package main

import "fmt"
import "net/http"
import "io"

func main() {
    fmt.Printf("server running...")

    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        io.WriteString(res, "Hello world from server")
    })

    http.ListenAndServe(":8080", nil)
}