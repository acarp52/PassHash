package main

import (
    //"os"
    "fmt"
    "log"
    "net/http"
    "time"
)

func routes(writer http.ResponseWriter, request *http.Request) {
    //request.ParseForm()

    if request.Method == "POST" {        
        request.ParseForm()
        
        if len(request.Form) != 1 {
            log.Println("Error: request contained multiple fields")
            fmt.Fprintf(writer, "Error: expected single field. Received %v\n", request.Form)
        } else {
            for key, _ := range request.Form {
                if key == "password" {
                    log.Println("Recieved password request, with body", request.Form[key][0])
                    hashedpasswd := hashPassword(request.Form[key][0])
                    log.Println(hashedpasswd)
                    showhash(writer,  hashedpasswd+ "\n")
                } else if key == "shutdown" {
                    fmt.Fprintf(writer, "You made a shutdown request!\n")
                    graceful()
                } else {
                    log.Printf("Recieved invalid request %v", key)
                    fmt.Fprintf(writer, "Server cannot handle your request, try  \"password\" or \"shutdown\".\n")
                }
            }
        }

        // params := request.Form["password"] // to get params value with key
        // fmt.Println(params)
        
        // log.Println(request.Form)
        // //LOG: map[{"test": "that"}:[]]
        // var t string
        // for key, _ := range request.Form {
        //     log.Println(len(request.Form[key]))
        //     log.Println(key)
        //     //LOG: {"test": "that"}
        //     // err := json.Unmarshal([]byte(key), &t)
        //     // if err != nil {
        //     //     log.Println(err.Error())
        //     // }
        // }
        // log.Println(t)
        // //LOG: that

        // if value, isValid := request.Form["password"]; isValid {
        //     password = hashPassword(value[0])
        // //} else if value, isValid := request.Form["shutdown"]; isValid {
        //    // password = "shutdown"
        // } else{
        //     // log error
        //     fmt.Fprintf(writer, "POST body not valid.\n")
        // }

        log.Printf("Hold open for 5 seconds")
        time.Sleep(time.Duration(5) * time.Second)
        //showhash(writer, hashPassword("angryMonkey") + "\n") 

    } else if request.Method == "GET" {
        inputpasswd(writer, nil) 
    } else {
        fmt.Fprintf(writer, "Invalid HTTP request. Try:\n$ curl -X POST --data 'password=[yourpassword]'' http://localhost:8080\n")
    }
}

func shutdown(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Server is shutting down and not accepting new requests. Sorry!\n")
}