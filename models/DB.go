package models

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const username = "root"
const password = "1234"
const host = "localhost"
const port = 3306
const database = "goCRUDbd"

func OpenDB() {
	connection, err := sql.Open("mysql", generateURL())
	if err != nil {
		panic(err)
	} else {
		db = connection
	}

}

func Ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database running")
	}
}

func CloseDatabase() {
	db.Close()
}

func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
}

// FUNCIONES CRUD

// INSERT USER
func InsertUser(user *User) error {
	sql := "INSERT INTO users(username, password, email) VALUES (?,?,?)"
	_, err := db.Exec(sql, user.Username, user.Password, user.Email)

	return err
}

// DELETE USER
func DeleteUser(id int) error {
	sql := "DELETE FROM users WHERE ID=?"
	_, err := db.Exec(sql, id)

	return err
}

// UPDATE USER
func UpdateUser(id int, user *User) error {
	sql := "UPDATE users SET username=?, password=?, email=? WHERE ID=?"
	_, err := db.Exec(sql, user.Username, user.Password, user.Email, id)

	return err
}

// GET USER
func GetUser(id int) (*User, error) {
	user := CreateUser()
	sql := "SELECT username, password, email FROM users WHERE ID=?"
	rows, _ := db.Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.Username, &user.Email, &user.Password)
	}

	if user.Username == "" && user.Password == "" && user.Email == "" {
		return nil, errors.New("the id that you send was incorrect")
	}

	return user, nil
}

// GET USERS
func GetUsers() []*User {
	users := []*User{}
	sql := "SELECT username, password, email FROM users"

	rows, _ := db.Query(sql)

	for rows.Next() {
		user := CreateUser()
		rows.Scan(&user.Username, &user.Password, &user.Email)

		users = append(users, user)
	}

	return users
}
