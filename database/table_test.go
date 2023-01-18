package database

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateTable(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()
	// script := "DROP TABLE user"
	script := "CREATE TABLE user(username VARCHAR(100) NOT NULL,password VARCHAR(100) NOT NULL, PRIMARY KEY(username))ENGINE = Innodb"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}
	fmt.Println("Succed Create table")

}

func TestTest(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()
	script := "INSERT INTO user(username,password) VALUES('admin', 'admin') "
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}
	fmt.Println("Succed Create table")

}

func TestLogin(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()

	username := "admin';#"
	password := "adasdsa"
	script := "SELECT username FROM user WHERE username ='" + username + "' AND password ='" + password + "'LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("succes login ", username)
	} else {
		fmt.Println("gagal login")
	}

}

func TestLoginQuerryParameter(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()

	username := "admin"
	password := "admin"
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("succes login ", username)
	} else {
		fmt.Println("gagal login")
	}

}

func TestInsertQuerrySafe(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()

	username := "eko1234"
	password := "eko"
	script := "INSERT INTO user(username,password) values(?,?)"
	fmt.Println(script)
	result, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		// handle error
	}
	fmt.Println("Last insert ID:", lastInsertID)

}
