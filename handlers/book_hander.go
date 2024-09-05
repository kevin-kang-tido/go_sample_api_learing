package handlers

import (
    "go_sample_api/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

type BookHandler struct {
    DB *gorm.DB
}

// Create a new book
// handleDataBaseConnection
func (handleDataBaseConnection *BookHandler) CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := handleDataBaseConnection.DB.Create(&book).Error;err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, book)
}

// Get a book by ID
func (h *BookHandler) GetBook(c *gin.Context) {
    var book models.Book
    if err := h.DB.Preload("Author").First(&book, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    c.JSON(http.StatusOK, book)
}

// Update a book by ID
func (h *BookHandler) UpdateBook(c *gin.Context) {
    var book models.Book
    if err := h.DB.First(&book, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    h.DB.Save(&book)
    c.JSON(http.StatusOK, book)
}

// Delete a book by ID
func (h *BookHandler) DeleteBook(c *gin.Context) {
    if err := h.DB.Delete(&models.Book{}, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
