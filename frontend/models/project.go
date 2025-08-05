package models

type Project struct {
	ProjectId uint    `json:"project_id" gorm:"column:Project_id;primaryKey;autoIncrement"`
	Title     string  `json:"title" gorm:"uniqueIndex;not null"`
	Desc      string  `json:"desc" gorm:"uniqueIndex;not null"`
	Tech      string  `json:"tech" gorm:"not null"`
	GitLink   *string `json:"git_link" gorm:"column:GitLink;default:null;"`
	WebLink   *string `json:"web_link" gorm:"column:WebLink;default:null;"`
	BlogId    *uint   `json:"blog_id" gorm:"column:Blog_id;default:null;"`
}
