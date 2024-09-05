package databases

import (
	"go_sample_api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB,error){
    
	// open a connection to the sqlite database 
	db, err := gorm.Open(sqlite.Open("book.db"),&gorm.Config{})
    

	if err != nil { // nil = zero values 
		return nil ,err; // Return error if fails to open database 
	}

    // AutoMigrate use to auto 
	if err := db.AutoMigrate(&models.Author{},&models.Book{});
	err != nil {
		return nil,err
	}

	return db,nil


}