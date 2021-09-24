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
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_mysql?parseTime=true")
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

/*
--------------------------------------Test querry sql complex--------------------------------------
*/

func TestQuerrySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //iterasi untuk menampilkan data didalam database
		var id, name, email string
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id", id, "name", name, "email", email, "balance", balance, "rating", rating, "birthdate", birthDate, "createdat", createdAt, "Married", married)

	}
}

/*
--------------------------------------SQL INJECTION--------------------------------------
*/

//Contoh Source Code yg rentan seranganan SQL Injection
//karena dengan  merubah script pada username akan membuat password tidak valid
func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin" //jika merubah username menjadi "admin'; #" maka akan menyebabkan querry setelah simbol # akan diabaikan
	password := "admin"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)

	}

	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Succes Login", username)
	} else {
		fmt.Println("Login Failed")
	}
}

/*
--------------------------------------SQL PARAMS--------------------------------------
*/

func TestSolveSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password) //tambahkan parameter yg akan disubtitusi pada script
	if err != nil {
		panic(err)

	}

	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Succes Login", username)
	} else {
		fmt.Println("Login Failed")
	}
}

/*
--------------------------------------Exec Context sql with params--------------------------------------
*/

func TestExecSqlWithParams(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "mock2" //Secure from sql injection
	password := "test"

	script := "INSERT INTO user(username, password) VALUE (?,?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done insert into user table")
}

/*
--------------------------------------Auto Increment--------------------------------------
*/

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "mock@gmail.com" //Secure from sql injection
	comments := "test comment"

	script := "INSERT INTO comments(email, comments) VALUE (?,?)"
	result, err := db.ExecContext(ctx, script, email, comments)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes insert comments with id", insertId)
}
