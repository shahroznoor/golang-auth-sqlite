package models

import (
	"errors"

	"auth.com/auth/db"
	"auth.com/auth/utils"
)

type User struct {
	ID       int64
	Name     string `binding:required`
	Email    string `binding:required`
	Password string `binding:required`
}

var users = []User{}

func (u *User) Save() error {
	query := `INSERT INTO user (name , email , password)
	VALUES (? , ? , ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Name, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	u.Password = hashedPassword

	return err

}

func GetAllUsers() ([]User, error) {
	query := `SELECT * FROM user`
	result, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var users []User
	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserById(id int64) (*User, error) {
	query := `SELECT * FROM user WHERE ID = ?`
	result := db.DB.QueryRow(query, id)
	var user User
	err := result.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id , password FROM user WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var retirevedPassword string
	err := row.Scan(&u.ID, &retirevedPassword)
	if err != nil {
		return errors.New("Invalid Email")
	}

	passwordIsValid := utils.CheckPassword(u.Password, retirevedPassword)
	if !passwordIsValid {
		return errors.New("Invalid Password")
	}
	return nil
}
