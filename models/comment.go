package models

// import "gorm.io/gorm"

// Comment represents a comment on a post.
type Comment struct {
    ID     uint   `gorm:"primaryKey" json:"id"`
    Userid uint   `json:"userId"`
    Title   string `json:"title"`
    Body   string `json:"body"`
}
