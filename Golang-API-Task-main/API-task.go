package main

import (
	"log"
    _ "github.com/lib/pq"
	"github.com/gin-gonic/gin"

)

func main() {
   
    db, err := initDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Gin router initialization
    router := gin.Default()

    // Middleware to pass the database connection to handlers
    router.Use(func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    })

    // Routes setup
    router.POST("/library", createLibrary)
    router.GET("/library", getAllLibrary)
    router.POST("/book/:email", createBook)
    router.GET("/book", getAllBook)
    router.PUT("/book/:isbn/:email", updateBookByISBN)
    router.DELETE("/book/:isbn/:email", deleteBookByISBN)
    router.POST("/user", createUser)
    router.PUT("/user/:id/:email", updateUserByID)
    router.GET("/book/:title", getBookByTitle)
    router.POST("/request", createRequest)
    router.GET("/request/:email", getAllRequest)
    router.PUT("/request/:reqid/:email", updateRequestByReqID)
    router.GET("/user", getAllUser)
    router.POST("/issue/:email", createIssue)
    // router.POST("/tasks", createTaskHandler)
    // router.GET("/tasks", getAllTasksHandler)
    // router.GET("/tasks/:id", getTaskByIDHandler)
    // router.PUT("/tasks/:id", updateTaskByID)
    // router.DELETE("/tasks/:id", deleteTaskByIDHandler)
    // Start the server
    router.Run(":6000")
}
