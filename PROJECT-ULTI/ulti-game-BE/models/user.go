package models

import (
	"errors"
	"fmt"

	"exmaple.com/ulti-restapi/database"
	"exmaple.com/ulti-restapi/utility"
)

type User struct {
	ID       int64
	Name     string
	Email    string `binding:"required"`
	Password string `binding:"required"`
	State    string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(name, email, password, state) VALUES(?,?,?,?)`
	stmt, err := database.Database.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashPassword, err := utility.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Name, u.Email, hashPassword, u.State)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = id
	return err
}
func (u *User) ValidateUser() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := database.Database.QueryRow(query, u.Email)

	var retrievesPass string

	err := row.Scan(&retrievesPass)
	fmt.Println(err)

	if err != nil {
		fmt.Println("ittttttt")
		return errors.New("credentials couldn't be read")
	}

	passWordisValid := utility.CheckPasswordHash(u.Password, retrievesPass)

	if !passWordisValid {
		return errors.New("credentials is invalid, password you entered not correct")
	}

	return nil

}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := database.Database.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.State)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
