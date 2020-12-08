package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Password string `json:"password"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Sex     string `json:"sex"`
}

func (b *User) TableName() string {
	return "user"
}
