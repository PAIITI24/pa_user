package model

type User struct {
	Id       int    `gorm:"primaryKey,autoIncrement" packets:"id"`
	Name     string `packets:"name"`
	Email    string `packets:"email" gorm:"unique"`
	Password string `packets:"password"`
	Role     int    `packets:"role"` // role 0 : owner, role 1 : employee/staff
}
