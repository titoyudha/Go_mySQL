package go_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDBConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost3306)/GO_MYSQL")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}

/*
--------------------------------------Database Pooling--------------------------------------
*/

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_mysql")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

/*
--------------------------------------SQL Execute --------------------------------------
*/

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUE ('user3', 'USER3')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done insert into customer")
}

/*
--------------------------------------SQL Querry --------------------------------------
*/

func TestQuerrySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //iterasi untuk menampilkan data didalam database
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id", id)
		fmt.Println("Name", name)
	}
}

/*
--------------------------------------Tipe Data Column--------------------------------------
*/

func TestExecSql2(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name, email, balance, rating, created_at, birth_date, married) VALUE ('user3','USER3','aaa@youremail.com', 2500000, 95.0,'22.10.10','2021-10-10', false);"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done insert into customer")
}
