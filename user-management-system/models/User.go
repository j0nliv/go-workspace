package models

type User struct {
	ID        int
	User      string
	Firstname string
	Lastname  string
	Profile   string
	Interests []Interest
}
