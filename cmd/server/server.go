package server

import (
    "fmt"
    "html"
    "log"
    "net/http"
)


func Server() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, nrm  %q", html.EscapeString(r.URL.Path))
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
