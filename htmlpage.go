package main

import "net/http"
import "fmt"

// Displays HTML if user wants to use their browser. In a production system,
// this should handle errors, but should be ok for the purposes of this demo  
func inputpasswd(w http.ResponseWriter) {
    htmlHead := `<html><head><title>passhash</title></head><body>`
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

// Displays the user's password hash. This could have easily been done in routes.go, but 
// I wanted to compartmentalize anything the user could see in a browser into this file.
func showhash(writer http.ResponseWriter, hashStr string) {
    fmt.Fprintf(writer, "%s", hashStr)
}
