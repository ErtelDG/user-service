package models

type User struct {
	UserID       int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	CreatedAt    string
	IsActive     bool
}