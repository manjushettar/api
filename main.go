package main

import (
    "fmt"
    "net/http"
    "api/handlers"
)

func main(){
    http.HandleFunc("/register", handlers.Register) 
    
    http.HandleFunc("/login", handlers.Login)

    http.HandleFunc("/getAll", handlers.GetAll)

    http.HandleFunc("/getByID", handlers.GetByID)
    
    http.HandleFunc("/getByEmail", handlers.GetByEmail)

    http.HandleFunc("/getByName", handlers.GetByName)
    
    http.HandleFunc("/getLoggedIn", handlers.GetAllLoggedIn)
    
    http.HandleFunc("/profile", handlers.AuthWrapper(handlers.GetProfile))

    fmt.Println("Server starting on port 8000")
    http.ListenAndServe(":8000", nil)
}
