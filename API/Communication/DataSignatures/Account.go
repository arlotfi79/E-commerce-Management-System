package DataSignatures

import (
	"time"
)

type Account struct {
	Id          uint64    `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	LastName    string    `json:"lastName" binding:"required"`
	UserName    string    `json:"username" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	PhoneNumber string    `json:"phoneNumber" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	BirthDate   time.Time `json:"birthDate" binding:"required"`
	JoinDate    time.Time `json:"JoinDate" binding:"required"`
}
