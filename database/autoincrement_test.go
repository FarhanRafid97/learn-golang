package database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestCreateTable2(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()
	// script := "DROP TABLE commenct"
	script := "CREATE TABLE comment(id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,email VARCHAR(100) NOT NULL,comment TEXT NOT NULL)ENGINE = Innodb"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}
	fmt.Println("Succed Create table")
}
func TestInsertAutoIncrement(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()
	email := "Farhan123@gmail.com"
	comment := "Test comment33 hehge"
	// script := "DROP TABLE user"
	script := "INSERT INTO comment(email,comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		// handle error
	}
	script2 := "SELECT id,email,comment from comment WHERE id = ?"
	rows, err := db.QueryContext(ctx, script2, lastInsertID)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var email, comment string
		err := rows.Scan(&id, &email, &comment)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID :", id)
		fmt.Println("Email :", email)
		fmt.Println("Comment :", comment)
	}

}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()
	script := "INSERT INTO comment(email,comment) values(?,?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		email := "farhan" + strconv.Itoa(i) + "@gmail.com"
		comment := "init comment ke " + strconv.Itoa(i)
		result, error := statement.ExecContext(ctx, email, comment)
		if error != nil {
			panic(error)
		}
		lastindex, errs := result.LastInsertId()
		if errs != nil {
			panic(errs)
		}
		fmt.Println("Ini id ke ", lastindex)
	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comment(email,comment) values(?,?)"

	for i := 30; i < 40; i++ {
		email := "farhan" + strconv.Itoa(i) + "@gmail.com"
		comment := "init comment ke " + strconv.Itoa(i)
		result, error := tx.ExecContext(ctx, script, email, comment)
		if error != nil {
			panic(error)
		}
		lastindex, errs := result.LastInsertId()
		if errs != nil {
			panic(errs)
		}
		fmt.Println("Ini id ke ", lastindex)
	}

	errs := tx.Commit()
	if errs != nil {
		panic(errs)
	}

}
