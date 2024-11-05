package main

import (
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"
    "io"
)

type post struct {
    userID string
    content string
}

type user struct {
    userID string
    name string
    email string
    password string
    posts []post
}

type registration struct {
    Name string `json:"name"`
    Password string `json:"password"`
    Email string `json:"email"`
}

var userId int = 0
var user_map map[string]user = make(map[string]user)

func parseRegistrationRequest(body io.ReadCloser) (*registration, error){
    decoder := json.NewDecoder(body)
    var reg registration
    err := decoder.Decode(&reg)

    defer body.Close()

    return &reg, err
}

func incrementUserIDOnNewUser() string {
    userId += 1
    return strconv.Itoa(userId)
}

func createNewUser(userID string, name string, email string, password string) *user{
    u := user{userID:userID, name:name, email:email, password: password}
    u.posts = make([]post, 0, 10)
    return &u
}

func register(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST"{
        reg, err := parseRegistrationRequest(r.Body)
        
        if err != nil{
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }
        
        name := reg.Name
        email := reg.Email
        password := reg.Password

        strRep := incrementUserIDOnNewUser()

        newUsr := createNewUser(strRep, name, email, password)
        user_map[strRep] = *newUsr 
        fmt.Fprintf(w, "User %v created.\n", name)
    } else{
        fmt.Fprintf(w, "Invalid request.\n")
    }
}


func getAll(w http.ResponseWriter, r *http.Request) {
    str := ""
    for _, usr := range user_map{
        str += "Name: " + usr.name + " email: " + usr.email + " id: " + usr.userID + "\n"
    }
    fmt.Fprintf(w, str)
}

func getByID(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET"{
        id := r.URL.Query().Get("id")

        val, ok := user_map[id]
        if !ok{
            fmt.Fprintf(w, "No user by that id\n")
        } else{
            fmt.Fprintf(w, "User: %v, email: %v, id: %v\n", val.name, val.email, val.userID)
        }
    }
}

func findUserByField(value string, field string) (*user, bool) {
    for _, val := range user_map{
        switch field{
        case "name":
            if val.name == value {
                return &val, true
            }
        case "email":
            if val.email == value {
                return &val, true
            }
        }
    }
    return &user{}, false
}

func getByName(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET"{
        name := r.URL.Query().Get("name")

        val, found := findUserByField(name, "name")
        if !found{
            fmt.Fprintf(w, "No user by that name\n")
        } else{
            fmt.Fprintf(w, "User: %v, email: %v, id: %v\n", val.name, val.email, val.userID)
        
        }
    }
}

func getByEmail(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET"{
        email := r.URL.Query().Get("email")

        val, found := findUserByField(email, "email")
        if !found{
            fmt.Fprintf(w, "No user by that email\n")
        } else{
            fmt.Fprintf(w, "User: %v, email: %v, id: %v\n", val.name, val.email, val.userID)
        
        }
    }
}

func main(){

    http.HandleFunc("/register", register) 
    
    http.HandleFunc("/getAll", getAll)

    http.HandleFunc("/getByID", getByID)
    
    http.HandleFunc("/getByEmail", getByEmail)

    http.HandleFunc("/getByName", getByName)

    fmt.Println("Server starting on port 8000")
    http.ListenAndServe(":8000", nil)
}
