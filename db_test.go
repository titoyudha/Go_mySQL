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
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost3306)/GO_MYSQL?parseTime=true")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}

/*
--------------------------------------Database Pooling--------------------------------------
*/

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost3306)/GO_MYSQL")
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
--------------------------------------SQL QUERRY--------------------------------------
*/

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer (id, name) VALUE ('user1', 'USER1')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done insert into customer")
}
