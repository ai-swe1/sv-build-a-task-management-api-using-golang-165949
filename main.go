package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    r := gin.Default()
    // Serve static assets
    r.Static("/static", "./static")

    // Register API routes
    InitRoutes(r)

    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}