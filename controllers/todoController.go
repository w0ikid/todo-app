package controllers

import (
    "net/http"
    "todo-app/database"
    "todo-app/models"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
}

func GetTodos(c *gin.Context) {
    session := sessions.Default(c)
    userID := session.Get("userID")

    var todos []models.Todo
    database.DB.Where("user_id = ?", userID).Find(&todos)
    c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
    var todo models.Todo
    session := sessions.Default(c)
    userID := session.Get("userID").(uint)
    todo.UserID = userID
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&todo)
    c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
    var todo models.Todo
    session := sessions.Default(c)
    userID := session.Get("userID").(uint)
    if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&todo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Save(&todo)
    c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
    var todo models.Todo
    session := sessions.Default(c)
    userID := session.Get("userID").(uint)
    if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&todo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    database.DB.Delete(&todo)
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
