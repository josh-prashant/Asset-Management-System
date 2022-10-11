package service

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Employee struct {
	EmpId     int    `json:"empId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Mobile    int    `json:"mobileNumber"`
}

type Asset struct {
	Id         int
	Name       string
	TotalCount int
	Available  int
	Allocate   int
	Type       string
	Expiry     string
	// Desc       string `json:"description"`
}

type Token struct {
	UserID uint
	Name   string
	Email  string
	Role   string
	*jwt.StandardClaims
}

type User struct {
	gorm.Model

	UserId   int
	Role     string
	Username string
	Password string `json:"Password"`
	EmpId    uint   `json:"EmpId"`
}

type AssetAlllocation struct {
	id      int
	EmpId   int
	Assetid int
}

func (AssetAlllocation) TableName() string {
	return "asset_allocation"
}

type Request struct {
	ReqId   int `gorm:"primaryKey"`
	Date    string
	Status  Status
	EmpId   int
	AssetId int
	Note    string
}

func (Request) TableName() string {
	return "request"
}
