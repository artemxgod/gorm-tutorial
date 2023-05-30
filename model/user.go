package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// we should specify index because we use it as a reference
	UserID int `gorm:"index;unique"`
	Name   string
	Age    int
	// foreign key let us store user_id in BankAccount.UserRefer
	// references overrides references from field ID to UserID
	BankAccount []*BankAccount `gorm:"foreignKey:UserRefer;references:UserID;constraint:OnDelete:CASCADE"`
}
