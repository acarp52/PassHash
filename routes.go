package main

import (
    "log"
    "net/http"
    "time"
)

// Controls how the server handles HTTP requests, and what the response is.
// Response headers are only changed in the event of an error, otherwise the
// default response code of 200 is sent along with the relevant response.
func routes(writer http.ResponseWriter, request *http.Request) {

    // Checks global variable to see if server is active and accepting new requests
    if serverActive{

        if request.Method == "POST" {        
            request.ParseForm()
            
            // Should only accept one parameter at a time
            if len(request.Form) != 1 {
                log.Println("Error: request contained multiple parameters")
                errmsg := "Error: expected single parameter." //" Received %v\n" + request.Form
                http.Error(writer, errmsg, 400)
            
            } else {
                // Parse requests
                for key, _ := range request.Form {
                    if key == "password" {
                        // Hash password and return it to the client
                        log.Println("Received password request, with body", request.Form[key][0])
                        hashedpasswd := hashPassword(request.Form[key][0])
                        log.Println(hashedpasswd)
                        showhash(writer,  hashedpasswd+ "\n")

                    } else if key == "shutdown" {
                        // Gracefully handle shutdown request.
                        // KNOW BUG: I couldn't figure out how to have the server complete its response
                        // before the program exits, so the client sees an empty response for this request.
                        errmsg := "Server shutdown initiated. No new requests can be made at this time."
                        http.Error(writer, errmsg, 503)

                        // Because this uses os.Exit() to cease operation, the HTTP response is never written.
                        // I tried multiple workarounds, but nothing worked.
                        graceful()

                    } else {
                        // Unknown request
                        log.Printf("Received invalid request %v", key)
                        errmsg := "Error: server cannot handle your request. Try  \"password\" or \"shutdown\"."
                        http.Error(writer, errmsg, 400)
                    }
                }
            }

            // Always hold valid connecton open for 5 seconds before responding
            log.Printf("Hold open for 5 seconds")
            time.Sleep(time.Duration(5) * time.Second)

        } else if request.Method == "GET" {
            // Display HTML. It will look strange if the user makes a GET through curl, but should 
            // display properly in a web browser.
            log.Println("Received GET")
            inputpasswd(writer) 
        
        } else {
            // User made an HTTP request that the server cannot handle.
            // Returns with an error message and HTTP response of: 501 Not Implemented
            log.Println("Received request method not supported by server.")
            errmsg := "Request method not supported by server. Try:\n$ curl -X POST --data 'password=[yourpassword]'' http://localhost:8080\n"
            http.Error(writer, errmsg, 501)

        }
    } else{
        // Reject any new requests if server is no longer active (in the process of shutting down)
        // Returns with an error message and HTTP response of: 503 Service Unavailable
        log.Println("Received new request during shutdown.")
        errmsg := "Server is shutting down and not accepting new requests. Sorry!"
        http.Error(writer, errmsg, 503)

    }
}

// Attempted workaround for shutting down.
func sutdownRequest(writer http.ResponseWriter, request *http.Request) {
    log.Println("something happenin...")
    errmsg := "Server shutdown initiated. No new requests can be made at this time."
    http.Error(writer, errmsg, 503)
}