package models

type Blog struct {
	BlogId uint   `json:"blog_id" gorm:"column:Blog_id;primaryKey;autoIncrement"`
	Title  string `json:"title" gorm:"column:Title;uniqueIndex;not null"`
	Date   string `json:"date" gorm:"column:Date;not null"`
}
