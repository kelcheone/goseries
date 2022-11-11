package models

type User struct {
	Id         int    `json:"id" gorm:"primaryKey; not null; autoIncrement"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Email      string `json:"email" gorm:"uniqueIndex"`
	Created_at int64  `json:"created_at" gorm:"autoCreateTime"`
}
