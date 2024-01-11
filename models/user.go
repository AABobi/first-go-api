package models

import (
	"errors"
	"example/db"
	"example/utils"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("save1")
		return err
	}

	defer stmt.Close()

	//Hash password
	fmt.Println(u.Password)
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		fmt.Println("save2")
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		fmt.Println("save3")
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	fmt.Println("save4")
	return err
}

func GetAllUser() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)

	fmt.Println("GetAllUser1")
	if err != nil {
		fmt.Println("GetAllUser2")
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		fmt.Println("rowsNext")
		var user User
		//W tym skanie znajduje sie pojedyńczy row
		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			fmt.Println("GetAllUser3")
			return nil, err
		}

		users = append(users, user)
	}
	fmt.Println("GetAllUser4", users)
	return users, nil

}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	//TO get a single row
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		fmt.Println("valid1l")
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		fmt.Println("valid2")
		return errors.New("Credentials invalid")
	}

	return nil
}
