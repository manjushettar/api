package utils

import (
    "golang.org/x/crypto/bcrypt"
    "encoding/json"
    "io"
    "strconv"
    "api/models"
)

func HashPassword(pass string) ([]byte, error) {
    return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword []byte, plainPassword string) error {
    return bcrypt.CompareHashAndPassword(hashedPassword, []byte(plainPassword))
}

func ParseRegistrationRequest(body io.ReadCloser) (*models.Registration, error) {
    decoder := json.NewDecoder(body)
    var reg models.Registration
    err := decoder.Decode(&reg)
    defer body.Close()
    return &reg, err
}

func ParseLoginRequest(body io.ReadCloser) (*models.LoginSession, error) {
    decoder := json.NewDecoder(body)
    var log models.LoginSession
    err := decoder.Decode(&log)
    defer body.Close()
    return &log, err
}

func IncrementUserID() string {
    models.UserID++
    return strconv.Itoa(models.UserID)
}

func GetAll() (string, bool){
    str := ""

    for _, val := range models.UserMap{
        str += "[" + val.UserID + "] " + val.Name + ": " + val.Email + "\n"
    }
    
    if len(str) == 0{
        return "", false
    }
    return str, true
}

func FindUserByField(value string, field string) (*models.User, bool) {
    for _, val := range models.UserMap {
        switch field {
        case "name":
            if val.Name == value {
                return &val, true
            }
        case "email":
            if val.Email == value {
                return &val, true
            }
        }
    }
    return &models.User{}, false
}

func FindAllUsersLoggedIn() (string, bool){
    str := ""

    for _, val := range models.UserMap {
        if val.LoggedIn {
            tmp :=  "[" + val.Email + "] " + val.Name + "\n"
            str += tmp
        }
    }
    
    if len(str) == 0{
        return "", false
    }

    return str, true
}

func IsLoggedIn(u *models.User) bool{
    if u.LoggedIn {
        return true 
    }
    return false
}

func LoginUser(u *models.User, p string) bool{
    err := VerifyPassword(u.Password, p)

    if err == nil{
        u.LoggedIn = true
        models.UserMap[u.UserID] = *u
        return true
    }

    return false
}

