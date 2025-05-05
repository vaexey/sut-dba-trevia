package model

type CommentReport struct {
	Id        uint `gorm:"primaryKey"`
	Content   string
	UserId    uint
	User      User `gorm:"foreignKey:UserId"`
	CommentId uint
	Comment   Comment `gorm:"foreignKey:CommentId"`
}