package models

type Post struct {
    UserID string
    Content string
}

type User struct {
    UserID string
    Name string
    Email string
    Password []byte
    LoggedIn bool
    Posts []Post
}

type Registration struct {
    Name string `json:"name"`
    Password string `json:"password"`
    Email string `json:"email"`
}

type LoginSession struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

var UserID int = 0
var UserMap map[string]User = make(map[string]User)

func CreateNewUser(userID string, name string, email string, hashedPassword []byte) (*User, error) {
    u := User{
        UserID:userID, 
        Name:name,
        Email:email,
        Password: hashedPassword,
        LoggedIn: false,
    }
    u.Posts = make([]Post, 0, 10)
    
    return &u, nil
}


