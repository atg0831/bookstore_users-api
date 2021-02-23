package users

import (
	"fmt"

	"github.com/atg0831/msabookstore/bookstore_users-api/datasources/mariadb/users_db"
	"github.com/atg0831/msabookstore/bookstore_users-api/utils/date_utils"
	"github.com/atg0831/msabookstore/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?,?,?,?);"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user id %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.ID = userID
	return nil
}
