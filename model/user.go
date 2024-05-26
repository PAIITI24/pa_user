package model

type User struct {
	Id       int    `json:"id" gorm:"primaryKey,autoIncrement" packets:"id"`
	Name     string `json:"name" packets:"name"`
	Email    string `json:"email" packets:"email" gorm:"unique"`
	Password string `json:"password" packets:"password"`
	Role     int    `json:"role" packets:"role"` // role 0 : owner, role 1 : employee/staff
}
