package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/article", showArticle)
    mux.HandleFunc("/article/create", createArticle)

    log.Println("Run server on http://127.0.0.1:4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
