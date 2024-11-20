package models

import (
	"errors"

	"shnk.com/eventx/db"
	"shnk.com/eventx/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `INSERT INTO users(email,password) VALUES(?,?)	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}
	res, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}
	id, err := res.LastInsertId()

	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (u *User) ValidateCredentials() error {
	var retrievedPassword string

	query := `SELECT password from users where email = ?`

	res := db.DB.QueryRow(query, u.Email)

	err := res.Scan(&retrievedPassword)
	if err != nil {
		return err
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !isPasswordValid {
		return errors.New("Invalid login credentials")
	}

	return nil
}
