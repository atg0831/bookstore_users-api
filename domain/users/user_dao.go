package users

import (
	"fmt"

	"github.com/atg0831/msabookstore/bookstore_users-api/utils/date"

	"github.com/atg0831/msabookstore/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
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
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("user email %s is already registered", user.Email))
		}

		return errors.NewBadRequestError(fmt.Sprintf("user id %d already exists", user.ID))
	}
	user.DateCreated = date.GetNowString()

	usersDB[user.ID] = user
	return nil
}
