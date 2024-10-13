package model

type User struct {
	*Model

	Email    string `json:"email"`
	Password string `json:"-"`
}

func (u *User) TableName() string {
	return "users"
}
