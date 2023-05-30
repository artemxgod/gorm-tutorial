package main

import (
	"fmt"
	"log"

	"github.com/artemxgod/gorm-tutorial/initial"
	"github.com/artemxgod/gorm-tutorial/queries"
)

var databaseURL = "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

func main() {

	// getting a database
	db, err := initial.GetDB(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// adding 2 user to a database
	if err := queries.CreateUsers(db); err != nil {
		log.Fatal(err)
	}

	// getting user Alex and check his bank accounts
	u, err := queries.GetUser(db, 101)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v %+v\n", u.BankAccount[0], u.BankAccount[1])

	// adding a bank account to user Michael
	_, err = queries.UpdateUserBankAccount(db, 102)
	if err != nil {
		log.Fatal(err)
	}

	// deleting user Alex and all of his accounts
	if err := queries.DeleteUser(db, 101); err != nil {
		log.Fatal(err)
	}

}
