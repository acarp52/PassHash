package main

import "net/http"
import "fmt"

func inputpasswd(w http.ResponseWriter, err error) {
    htmlHead := `<html><body>`
    htmlForm := 
        `<form action="/" method="post">Enter Password:<br>
            <input type="text" name="password">
            <input type="submit" value="Hash">
        </form>`

    htmlEnd := `</body></html>`

    fmt.Fprintf(w, "%s", htmlHead)
    fmt.Fprintf(w, "%s", htmlForm)
    fmt.Fprintf(w, "%s", htmlEnd)
}

func showhash(writer http.ResponseWriter, hashStr string) {
    fmt.Fprintf(writer, "%s", hashStr)
}
