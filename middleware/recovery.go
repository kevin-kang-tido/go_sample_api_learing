package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "log"
)

func RecoveryMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Recovered from panic: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
                c.Abort()
            }
        }()
        c.Next()
    }
}
