package models

// import "gorm.io/gorm"

// Comment represents a comment on a post.
type Comment struct {
    ID     uint   `gorm:"primaryKey" json:"id"`
    PostID uint   `json:"post_id"`
    Name   string `json:"name"`
    Email  string `json:"email"`
    Body   string `json:"body"`
}
