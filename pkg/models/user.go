package models

// User represents a user
// @Description user model
// @receiver u
// @return User
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}