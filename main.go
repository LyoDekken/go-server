package main

import (
    "net/http"

    "goServer/api/routes"
)

func main() {
    http.HandleFunc("/hello", routes.HelloHandler)

    http.ListenAndServe(":8080", nil)
}
