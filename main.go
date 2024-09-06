package main

import (
  
	"go_sample_api/databases"
	"go_sample_api/handlers"
	"github.com/gin-gonic/gin"
  "github.com/swaggo/files"
  "github.com/swaggo/gin-swagger"
  // _ "go_sample_api/docs" // Import the generated docs
)

func main() {

    // inittainlize  dastabase 
    db ,err := databases.InitDB()
    if err != nil {
        // panic immediately to the norml flow 
        panic("Fial to connect to the database")
    }

    // handler 
    authorHandler := &handlers.AuthorHander{DB: db}
    bookHandler := &handlers.BookHandler{DB:db}

    // defualt gin router 
    router := gin.Default();

    // Author routes
    router.POST("/authors", authorHandler.CreateAuthor)
    router.GET("/authors/:id", authorHandler.GetAuthor)
    router.PUT("/authors/:id", authorHandler.UpdateAuthor)
    router.DELETE("/authors/:id", authorHandler.DeteleAuthor)

    // books route
    // Book routes
    router.POST("/books",bookHandler.CreateBook)
    router.GET("/books/:id", bookHandler.GetBook)
    router.PUT("/books/:id", bookHandler.UpdateBook)
    router.DELETE("/books/:id", bookHandler.DeleteBook)

    // Swagger UI route
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  router.Run();
}

