package model

type User struct {
	UserID   int    `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
}
