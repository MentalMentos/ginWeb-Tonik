package model

type User struct {
	ID    int64  `gorm:"type:int; primary_key" json:"id"`
	Name  string `gorm:"type:string" json:"name"`
	Email string `gorm:"type:string" json:"email"`
}
