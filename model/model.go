package model

type User struct {
	Id      int    `json:"-"`
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Login   string `json:"login" binding:"required"`
	Pass    string `json:"pass" binding:"required"`
	Email   string `json:"email" binding:"required"`
}
