package utils

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    "encoding/json"
    "io"
    "strconv"
    "api/models"
    "time"
    "errors"
)

type Claims struct {
    UserID string `json:"user_id"`
    Email string `json:"email"`
    jwt.RegisteredClaims
}

func CreateToken(userID string, email string) (string, error){
    expirationTime := time.Now().Add(1*time.Hour)

    claims := Claims{
        UserID: userID,
        Email: email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
            IssuedAt: jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    secretKey := []byte("key")
    tokenString, err := token.SignedString(secretKey)
    return tokenString, err
}

func VerifyToken(tokenString string) (*Claims, error){
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func ( token *jwt.Token) (interface{}, error){
        return []byte("key"), nil
    })

    if err != nil{
        return claims, nil
    }
    
    if !token.Valid{
        return nil, errors.New("Invalid token")
    }

    return claims, nil
}

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

