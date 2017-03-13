package main

import (
    "os"
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    if len(os.Args) != 2 {
        errOut()
        return
    }

    port := os.Args[1]
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        log.Println(r.Form)
        fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))

}

func errOut() {
    fmt.Println("Usage: \n")
}
