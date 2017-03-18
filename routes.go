package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

// Countrols how the server handles HTTP requests, and what the response is
func routes(writer http.ResponseWriter, request *http.Request) {

    // Checks global variable to see if server is active and accepting new requests
    if serverActive{

        if request.Method == "POST" {        
            request.ParseForm()
            
            // Should only accept one parameter at a time
            if len(request.Form) != 1 {
                log.Println("Error: request contained multiple parameters")
                fmt.Fprintf(writer, "Error: expected single parameter. Received %v\n", request.Form)
            
            } else {
                // Parse requests
                for key, _ := range request.Form {
                    if key == "password" {
                        log.Println("Recieved password request, with body", request.Form[key][0])
                        hashedpasswd := hashPassword(request.Form[key][0])
                        log.Println(hashedpasswd)
                        showhash(writer,  hashedpasswd+ "\n")

                    } else if key == "shutdown" {
                        // fmt.Fprintf(writer, "You made a shutdown request!\n")
                        errmsg := "Server shutdown initiated. No new requests can be made at this time."
                        http.Error(writer, errmsg, 503)
                        graceful()

                    } else {
                        log.Printf("Recieved invalid request %v", key)
                        errmsg := "Error: server cannot handle your request. Try  \"password\" or \"shutdown\"."
                        http.Error(writer, errmsg, 500)
                    }
                }
            }

            // Always hold valid connecton open for 5 seconds before responding
            log.Printf("Hold open for 5 seconds")
            time.Sleep(time.Duration(5) * time.Second)

        } else if request.Method == "GET" {
            // Display HTML. It will look strange if the user makes a GET through curl, but should 
            // display properly in a web browser.
            log.Prinln("Recieved GET")
            inputpasswd(writer) 
        
        } else {
            // User made an HTTP request that the server cannot handle.
            // Returns with an error message and HTTP response of: 501 Not Implemented
            errmsg := "Request method not supported by server. Try:\n$ curl -X POST --data 'password=[yourpassword]'' http://localhost:8080\n"
            http.Error(writer, errmsg, 501)

        }
    } else{
        // Reject any new requests if server is no longer active (in the process of shutting down)
        // Returns with an error message and HTTP response of: 503 Service Unavailable
        errmsg := "Server is shutting down and not accepting new requests. Sorry!"
        http.Error(writer, errmsg, 503)
    }
}