package models

import (
	"errors"
	"event-booking-rest-api/db"
	"event-booking-rest-api/utils"
)

type User struct {
	ID            int64
	User_name     string `binding: "required"`
	User_surname  string `binding: "required"`
	User_email    string `binding: "required"`
	User_password string `binding: "required"`
}

func (u User) Save() error{

	query := ` INSERT INTO Users(user_name, user_surname, user_email, user_password)
	VALUES(?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil{
		return err
	}

	defer stmt.Close()

	hashpass, err := utils.HashPassword(u.User_password)

	if err != nil{
		return err
	}

	user, err := stmt.Exec(u.User_name, u.User_surname, u.User_email, hashpass)

	if err != nil{
		return err
	}
 
	id, err := user.LastInsertId()
	u.ID = id
	return err
}


func (u *User)ValidateUser () error {
	query := `SELECT ID, user_password FROM Users where user_email = ?`
	row := db.DB.QueryRow(query, u.User_email)

	var retrivedPassword string	

	err := row.Scan(&u.ID, &retrivedPassword)

	if err != nil{
		return errors.New("Credentials are Invalid")
		// return

	}


	passwordIsValid := utils.CheckPassword(u.User_password,  retrivedPassword)

	if !passwordIsValid{
		return errors.New("Credentials are Invalid")
	}


	return nil
}