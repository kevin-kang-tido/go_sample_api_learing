package handlers

import (
	"go_sample_api/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthorHander struct {
	DB *gorm.DB;
}

// handle create new author to database 
func (handleCreateAuthour * AuthorHander) CreateAuthor(create *gin.Context){
	
	var author models.Author
    
	if err := create.ShouldBindJSON(&author);

	err != nil { // err = !nil * it is mean(if have error )
	create.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
    return
	}

	if err := handleCreateAuthour.DB.Create(&author).Error; // .Error  part is used to check if an error occurred during the execution of the database operation
	err != nil{
		create.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		return
	}
	
	create.JSON(http.StatusOK,author)
}

// get by id 
func (handleGetAuthor *AuthorHander) GetAuthor(getRequest *gin.Context){
    
	// load author from database
	var author models.Author

	if err := handleGetAuthor.DB.Preload("Books").First(&author,getRequest.Param("id")).Error;

	err !=nil{
		getRequest.JSON(http.StatusNotFound,gin.H{"error": "Author has been not found"})
	}

	getRequest.JSON(http.StatusOK,author)

}

// handle update author 

func (handleUpdateAuthor *AuthorHander) UpdateAuthor(getRequest *gin.Context){
    
	var author models.Author // declare a varible authour that is have type from models.author 

	if err := handleUpdateAuthor.DB.First(&author , getRequest.Param("id")).Error;
	err != nil{
		getRequest.JSON(http.StatusNotFound, gin.H{"erorr" : "Author has been not found!"})
		return
	}

	if err := getRequest.ShouldBindJSON(&author);err != nil{
		getRequest.JSON(http.StatusBadRequest, gin.H{"error":"Please check your json structure agian!"})
		return
	}

	handleUpdateAuthor.DB.Save(&author)
	getRequest.JSON(http.StatusOK, gin.H{
		"message":"Author updated sueccessfully!",
		"author":author,
	})

}


// detete author by id 
func(handleDeteleAuthor *AuthorHander) DeteleAuthor(getRequest *gin.Context){

	if err := handleDeteleAuthor.DB.Delete(&models.Author{}, getRequest.Param("id")).Error;
	err != nil{
		getRequest.JSON(http.StatusInternalServerError, gin.H{"eror": err.Error()})
	 return
	}
	
	getRequest.JSON(http.StatusOK, gin.H{"message":"Author has been deleted successfuly !"})
}




