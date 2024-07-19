package middleware

import (
    "net/http"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "fmt"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("AuthMiddleware executed")
        session := sessions.Default(c)
        userID := session.Get("userID")
        if userID == nil {
            fmt.Println("Unauthorized access")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        fmt.Println("User authorized:", userID)
        c.Next()
    }
}
