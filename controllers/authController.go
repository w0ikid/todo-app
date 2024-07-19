package controllers

import (
    "net/http"
    "todo-app/database"
    "todo-app/models"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Register(c *gin.Context) {
    var creds Credentials
    if err := c.BindJSON(&creds); err != nil {
        fmt.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("Error generating hashed password:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    user := models.User{Username: creds.Username, Password: string(hashedPassword)}
    result := database.DB.Create(&user)
    if result.Error != nil {
        fmt.Println("Error creating user:", result.Error)
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
        return
    }

    fmt.Println("User registered:", user.Username)

    c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
    var creds Credentials
    if err := c.BindJSON(&creds); err != nil {
        fmt.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := database.DB.Where("username = ?", creds.Username).First(&user).Error; err != nil {
        fmt.Println("User not found:", creds.Username)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
        fmt.Println("Invalid password for user:", creds.Username)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    session := sessions.Default(c)
    session.Set("userID", user.ID)
    session.Save()

    fmt.Println("User logged in:", user.Username)
    c.Redirect(http.StatusFound, "/")
}

func Logout(c *gin.Context) {
    session := sessions.Default(c)
    session.Clear()
    session.Save()
    c.Redirect(http.StatusFound, "/login")
}
