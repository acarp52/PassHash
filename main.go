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

func usage() {
        fmt.Fprintf(os.Stderr, "Usage: passhash [port number]\n")
        os.Exit(1)
}

func graceful() {
    //mux := http.NewServeMux()
    http.HandleFunc("/", shutdown)
    // http.ListenAndServe(":8080", mux)

    fmt.Println("Completing open requests and shutting down server in 5 seconds...")
    time.Sleep(time.Duration(5) * time.Second)
    fmt.Println("Shutdown complete. Goodbye!")
    os.Exit(1)
}

func main() {
    // Make sure a port number is supplied
    if len(os.Args) != 2 {
        usage()
    }
    port := os.Args[1]
    mux := http.NewServeMux()

    // Gracefully handle CTRL-C interrupt from the OS
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        graceful()
    }()

    for {
        
        mux.HandleFunc("/", routes) 
        // log.Fatal(http.ListenAndServe(":" + port, nil))
        log.Printf("Server is listening at port %v", port)
        http.ListenAndServe(":" + port, mux)
    }
}
