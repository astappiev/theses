package user

type LoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
