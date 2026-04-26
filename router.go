package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description,omitempty"`
    DueDate     string `json:"due_date,omitempty"` // ISO‑8601 date string
    Completed   bool   `json:"completed"`
}

// Global DB placeholder – replace with real SQLite handling.
var db = InitDB()

func InitRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.GET("/tasks", getTasks)
        api.POST("/tasks", createTask)
        api.GET("/tasks/:id", getTask)
        api.PUT("/tasks/:id", updateTask)
        api.DELETE("/tasks/:id", deleteTask)
        api.PATCH("/tasks/:id/complete", completeTask)
    }
}

func getTasks(c *gin.Context) {
    // Optional filter: ?completed=true|false
    filter := c.Query("completed")
    tasks, err := dbGetTasks(filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
    var t Task
    if err := c.ShouldBindJSON(&t); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    t.Completed = false
    id, err := dbCreateTask(t)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    t.ID = id
    c.JSON(http.StatusCreated, t)
}

func getTask(c *gin.Context) {
    // Implementation left as an exercise – retrieve by ID
    c.Status(http.StatusNotImplemented)
}

func updateTask(c *gin.Context) {
    // Implementation left as an exercise – full update by ID
    c.Status(http.StatusNotImplemented)
}

func deleteTask(c *gin.Context) {
    // Implementation left as an exercise – delete by ID
    c.Status(http.StatusNotImplemented)
}

func completeTask(c *gin.Context) {
    // Implementation left as an exercise – set Completed = true
    c.Status(http.StatusNotImplemented)
}

// ----- Minimal SQLite helpers (placeholders) -----

func InitDB() *DB {
    // In a real project, open the SQLite file, run schema.sql if needed.
    // This stub returns an empty struct to keep the compiler happy.
    return &DB{}
}

type DB struct{}

func dbGetTasks(filter string) ([]Task, error) {
    // Stub: return empty slice. Replace with actual SELECT query.
    return []Task{}, nil
}

func dbCreateTask(t Task) (int, error) {
    // Stub: return dummy ID. Replace with INSERT statement.
    return 1, nil
}