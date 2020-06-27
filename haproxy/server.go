package main

import (
    "flag"
    "fmt"
    "net/http"
)

func indexHandler(arg string) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "hello world!!!!" + arg)
    }
}

func main() {
    arg := flag.String("who", "1111", "an args to mark who use it.")
    flag.Parse()
    http.HandleFunc("/", indexHandler(*arg))
    http.ListenAndServe(":8000", nil)
}
