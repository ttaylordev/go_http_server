package main

import "net/http"
import "log"
import "os"

func main() {
    logger := log.New(os.Stdout, "", log.LstdFlags)
    logger.Println("Server started on http://localhost:8082")

     http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request){
        
        _, err := w.Write([]byte("Hello world"))
        if err != nil {

            // Handle the error
            logger.Println("Error writing response: ", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }
    })

    err := http.ListenAndServe(":8082", nil)
    if err != nil{
        logger.Fatal("Error starting server: ", err)
    }
    
}
