package main

import (
    "todo-app/controllers"
    "todo-app/database"
    "todo-app/middleware"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    store := cookie.NewStore([]byte("secret"))
    r.Use(sessions.Sessions("mysession", store))
    r.LoadHTMLGlob("templates/*")
    r.Static("/static", "./static")

    database.SetupDatabase()

    r.GET("/register", func(c *gin.Context) {
        c.HTML(http.StatusOK, "register.html", nil)
    })

    r.GET("/login", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", nil)
    })

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.POST("/logout", controllers.Logout)

    // middleware AuthMiddleware
    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.GET("/", controllers.ShowIndexPage)
        auth.GET("/todos", controllers.GetTodos)
        auth.POST("/todos", controllers.CreateTodo)
        auth.PUT("/todos/:id", controllers.UpdateTodo)
        auth.DELETE("/todos/:id", controllers.DeleteTodo)
    }

    r.Run()
}
