package DataSignatures

import (
	"time"
)

type Account struct {
	NationalCode string    `json:"nationalCode" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	LastName     string    `json:"lastName" binding:"required"`
	UserName     string    `json:"username" binding:"required"`
	Password     string    `json:"password" binding:"required"`
	PhoneNumber  string    `json:"phoneNumber" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	Gender       bool      `json:"gender" binding:"required"`
	BirthDate    time.Time `json:"birthDate" binding:"required"`
	JoinDate     time.Time `json:"JoinDate" binding:"required"`
}
