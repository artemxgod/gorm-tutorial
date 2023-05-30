package queries

import (
	"github.com/artemxgod/gorm-tutorial/model"
	"gorm.io/gorm"
)

func CreateUsers(db *gorm.DB) error {
	if err := db.Create(&model.User{
		UserID: 101,
		Name:   "Alex",
		Age:    100,
		BankAccount: []*model.BankAccount{
			{
				Money:    10000,
				Currency: "USD",
			},
			{
				Money:    5000,
				Currency: "EUR",
			},
		},
	}).Error; err != nil {
		return err
	}

	if err := db.Create(&model.User{
		UserID: 102,
		Name:   "Michael",
		Age:    100,
	}).Error; err != nil {
		return err
	}

	return nil
}

func GetUser(db *gorm.DB, id int) (*model.User, error) {
	user := &model.User{}
	// DB.Preload() retrives information from band_accounts table
	if err := db.Where(&model.User{UserID: id}).Preload("BankAccount").Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserBankAccount(db *gorm.DB, id int) (*model.User, error) {
	user, err := GetUser(db, id)
	if err != nil {
		return nil, err
	}

	user.BankAccount = append(user.BankAccount, &model.BankAccount{
		Money:    15000,
		Currency: "CAD",
	})

	user.Age = 101

	if err := db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(db *gorm.DB, id int) error {
	// soft delete
	err := db.Where(&model.User{UserID: id}).Delete(&model.User{}).Error
	if err != nil {
		return err
	}

	// permanent delete
	db.Unscoped().Delete(&model.User{}, 1)


	return nil
}
