package models

type User struct {
	Id string
	Name string
}

func NewUser(name string) (*User, error) {
	return &User{Id: "id", Name: name}, nil
}