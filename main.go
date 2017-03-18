package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
    "os/signal"
    "syscall"
    "time"
)

// Global variable that should be true while the server is active and accepting new requests
var serverActive bool

func main() {

    // Make sure a port number is supplied
    if len(os.Args) != 2 {
        usage()
    }
    port := os.Args[1]
    
    // Gracefully handle CTRL-C interrupt from the OS
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        graceful()
    }()

    // Start server
    serverActive = true
    http.HandleFunc("/", routes) 
    log.Printf("Server is listening at port %v", port)
    http.ListenAndServe(":" + port, nil)
}

// Outputs usage information and exits the program
func usage() {
        fmt.Fprintf(os.Stderr, "Usage: passhash [port number]\n")
        os.Exit(1)
}

// Gracefully shuts down the server. 
// See routes.go for handling of the HTTP requests during this process
func graceful() {
    serverActive = false
    log.Println("Server shutdown initiated. Completing open requests and shutting down in 5 seconds...")
    time.Sleep(time.Duration(5) * time.Second)
    log.Println("Shutdown complete. Goodbye!")
    os.Exit(1)
}

