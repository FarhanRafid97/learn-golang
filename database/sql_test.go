package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestSqlExec(t *testing.T) {

	db := GetConnection()

	ctx := context.Background()
	defer db.Close()
	script := "INSERT INTO mytable(name) VALUES('farhan12')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("seccess added new data")

}

type User struct {
	id                     int
	name                   string
	email                  string
	balance                int
	rating                 float32
	created_at, birth_date time.Time
	married                bool
}

func TestGetData(t *testing.T) {
	arr := make([]User, 0)
	db := GetConnection()

	ctx := context.Background()
	defer db.Close()
	script := "SELECT id, name FROM mytable where id > 5"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		arr = append(arr, User{id: id, name: "Name : " + name})

	}
	fmt.Println(arr)
	defer rows.Close()
	fmt.Println("seccess added new data")

}

func TestInsertdata(t *testing.T) {
	arr := make([]User, 0)
	db := GetConnection()

	ctx := context.Background()
	defer db.Close()
	script := "INSERT INTO mytable(name) VALUES('farhan322')"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		arr = append(arr, User{id: id, name: "Name : " + name})

	}
	fmt.Println(arr)
	defer rows.Close()
	fmt.Println("seccess added new data")

}

func TestQuerry(t *testing.T) {
	arr := make([]User, 0)
	db := GetConnection()

	// "INSERT INTO mytable(name,email,balance,rating,birth_date,married) VALUES('farhan','farhan@gmail.com',100000,4.5,'1997-10-10',1)"
	defer db.Close()
	ctx := context.Background()
	script := "SELECT id, name, email, balance, rating, married,created_at,birth_date married FROM mytable"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name, email string
		var married bool
		var rating float32
		var balance int
		var created_at, birth_date time.Time

		err := rows.Scan(&id, &name, &email, &balance, &rating, &married, &created_at, &birth_date)
		if err != nil {
			panic(err)
		}
		fmt.Println(created_at)
		fmt.Println(balance)
		fmt.Println(rating)
		arr = append(arr, User{id: id, name: name, email: email, married: married, created_at: created_at, birth_date: birth_date})

	}
	rows.Close()
	fmt.Println(arr)

}
