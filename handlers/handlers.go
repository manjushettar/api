package handlers

import (
    "net/http"
    "fmt"
    "api/models"
    "api/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        fmt.Fprintf(w, "Invalid request.\n")
        return
    }

    reg, err := utils.ParseRegistrationRequest(r.Body)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    _, exists := utils.FindUserByField(reg.Email, "email")
    if exists {
        http.Error(w, "Email already taken.", http.StatusBadRequest)
        return
    }

    hashedPassword, err := utils.HashPassword(reg.Password)
    if err != nil {
        http.Error(w, "Error processing password", http.StatusInternalServerError)
        return
    }

    userID := utils.IncrementUserID()
    newUser, err := models.CreateNewUser(userID, reg.Name, reg.Email, hashedPassword)
    
    if err != nil{
        http.Error(w, "Error creating new user.", http.StatusInternalServerError)
        return
    }

    models.UserMap[userID] = *newUser

    fmt.Fprintf(w, "User %v created.\n", reg.Name)
}

func Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST"{
        fmt.Fprintf(w, "Invalid request.\n")
        return 
    }
    
    log, err := utils.ParseLoginRequest(r.Body)

    if err != nil{
        http.Error(w, "Invalid login request", http.StatusBadRequest)
        return
    }

    usr, exists := utils.FindUserByField(log.Email, "email")
    if !exists {
        http.Error(w, "Email not found", http.StatusBadRequest)
        return
    }
    
    loggedIn := utils.IsLoggedIn(usr)
    if loggedIn{
        http.Error(w, "Already logged in", http.StatusBadRequest)
        return
    }

    ok := utils.LoginUser(usr, log.Password)
    if !ok {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    } else{
        fmt.Fprintf(w, "[%v] logged in\n", log.Email) 
    }
}

func GetAll(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET"{
        fmt.Fprintf(w, "Invalid Request.\n")
        return
    }
    
    str, ok := utils.GetAll()
    
    if !ok{
        fmt.Fprintf(w, "No users.\n")
        return
    }

    fmt.Fprintf(w, str)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET"{
        fmt.Fprintf(w, "Invalid request.\n")
        return
    }
    
    id := r.URL.Query().Get("id")

    val, ok := models.UserMap[id]
    if !ok{
        fmt.Fprintf(w, "No user by that id\n")
        return
    } 
    
    fmt.Fprintf(w, "User: %v, email: %v, id: %v\n", val.Name, val.Email, val.UserID)
}


func GetByName(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET"{
        fmt.Fprintf(w, "Invalid request.\n")
        return
    }
    
    name := r.URL.Query().Get("name")

    val, found := utils.FindUserByField(name, "name")
    if !found{
        fmt.Fprintf(w, "No user by that name\n")
        return
    }
    
    fmt.Fprintf(w, "User: %v, email: %v, id: %v\n", val.Name, val.Email, val.UserID)
}

func GetByEmail(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET"{
        fmt.Fprintf(w, "Invalid request.\n")
        return
    }

    email := r.URL.Query().Get("email")

    val, found := utils.FindUserByField(email, "email")
    
    if !found{
        fmt.Fprintf(w, "No user by that email\n")
        return
    } 
    
    fmt.Fprintf(w, "User: %v, email: %v, id: %v\n", val.Name, val.Email, val.UserID)
}

func GetAllLoggedIn(w http.ResponseWriter, r *http.Request){
    if r.Method != "GET"{
        fmt.Fprintf(w, "Invalid request.\n")
        return
    }
    
    val, found := utils.FindAllUsersLoggedIn()

    if !found {
        fmt.Fprintf(w, "No logged in users\n")
        return
    }
    
    fmt.Fprintf(w, val)
}


