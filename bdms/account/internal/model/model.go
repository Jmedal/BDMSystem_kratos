package model

import "github.com/bilibili/kratos/pkg/time"

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

type Article struct {
	ID      int64
	Content string
	Author  string
}

// User for sobot userinfo
type User struct {
	ID         int64     `json:"userId"`
	Avatar     string    `json:"avatar"`
	Account    string    `json:"account"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	Birthday   time.Time `json:"birthday"`
	Sex        string    `json:"sex"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	DeptId     string    `json:"deptId"`
	Status     int32     `json:"status"`
	CreateTime time.Time `json:"createTime"`
}
