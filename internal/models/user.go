package models

type User struct {
	id   int
	name string
}

func (u *User) CreateNewUser(Name string) *User {
	ID++
	return &User{id: ID, name: Name}
}

func (u *User) GetUserInfo() (int, string) {
	return u.id, u.name
}

var UserTemplate = User{id: 0, name: "[]"}
var ID = 0
