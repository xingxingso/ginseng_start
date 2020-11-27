package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	SexAge  string `json:"sex_age"`
	Sex     string `json:"sex"`
}

func (b *User) TableName() string {
	return "user"
}
