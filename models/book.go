package models

type Book struct {
    ID       uint   `gorm:"primaryKey" json:"id"`
    Title    string `json:"title"`
    AuthorID uint   `json:"author_id"`
    Author   Author `json:"author" gorm:"foreignKey:AuthorID"`
}
