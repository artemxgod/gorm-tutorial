package model

import "gorm.io/gorm"

type BankAccount struct {
	gorm.Model
	Money     int
	Currency  string
	UserRefer int
}
