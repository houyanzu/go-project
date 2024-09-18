package user

import (
	"github.com/houyanzu/work-box/database"
	"time"
)

type User struct {
	ID         uint   `json:"id" gorm:"id"`
	Email      string `json:"email" gorm:"email"`
	Phone      string `json:"phone" gorm:"phone"`
	Password   string `json:"password" gorm:"password"`
	CreateTime int64  `json:"create_time" gorm:"create_time"`
}

func (*User) TableName() string {
	return "user"
}

type Model struct {
	*database.MysqlContext
	Data  User
	List  []User
	Total int64
}

func New(ctx *database.MysqlContext) *Model {
	if ctx == nil {
		ctx = database.GetContext()
	}
	list := make([]User, 0)
	data := User{}

	return &Model{ctx, data, list, 0}
}

func (m *Model) Add() {
	m.Data.CreateTime = time.Now().Unix()
	m.Db.Create(&m.Data)
}

func (m *Model) InitByID(ID uint) *Model {
	m.Db.Take(&m.Data, ID)
	return m
}
