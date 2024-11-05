package main

import (
    "fmt"
    "net/http"
    "strconv"
)

type post struct {
    userID string
    content string
}

type user struct {
    userID string
    name string
    email string
    posts []post
}

func createNewUser(userID string, name string, email string) *user{
    u := user{userID:userID, name:name, email:email}
    u.posts = make([]post, 0, 10)
    return &u
}

func main(){
    userId := 0
    user_map := make(map[string]user)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello World!")
    })
    
    http.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request){
        if r.Method == "GET"{
            name := r.URL.Query().Get("name")
            email := r.URL.Query().Get("email")


            curId := userId + 1
            strRep := strconv.Itoa(curId)
            
            userId += 1
          
            newUsr := createNewUser(strRep, name, email)
            user_map[strRep] = *newUsr 
            fmt.Fprintf(w, "User %v created.\n", name)
        } else{
            fmt.Fprintf(w, "Invalid request.\n")
        }
    }) 
    

    http.HandleFunc("/getAll", func(w http.ResponseWriter, r *http.Request){
        str := ""
        for id, usr := range user_map{
            str += "Name: " + usr.name + " email: " + usr.email + " id: " + usr.id + "\n"
        }
        fmt.Fprintf(w, str)
    })

    http.HandleFunc("/getUser", func(w http.ResponseWriter, r *http.Request){
        if r.Method == "GET"{
            id := r.URL.Query().Get("id")

            val, ok := user_map[id]
            if !ok{
                fmt.Fprintf(w, "No user by that name\n")
            } else{
                fmt.Fprintf(w, "User: %v, email: %v, id: %v\n", val.name, val.email, val.userID)
            }
        }
    })

    fmt.Println("Server starting on port 8000")
    http.ListenAndServe(":8000", nil)
}
