package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })
    dt := time.Now()
      
    // printing the time in string format
    fmt.Println("Current date and time is: ", dt.String())
    http.ListenAndServe(":8080", nil)
}
