package user

type User struct {
	Id        uint `gorm:"primaryKey"`
	Email     string
	Title     string
	FirstName string
	LastName  string
	FullName  string
}

type LoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
