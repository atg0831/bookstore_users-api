package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/atg0831/msabookstore/bookstore_users-api/constant"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB

	port     string = os.Getenv(constant.MysqlUsersPort)
	schema   string = os.Getenv(constant.MysqlUsersSchema)
	host     string = os.Getenv(constant.MysqlUsersHost)
	password string = os.Getenv(constant.MysqlUsersPassword)
	username string = os.Getenv(constant.MysqlUsersUsername)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		// username, password, host, schema,
		"root",            //username
		"1234",            //password
		"127.0.0.1:33066", //host
		"users_db",        //schema
	)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
