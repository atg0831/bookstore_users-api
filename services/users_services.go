package services

import (
	"strings"

	"github.com/atg0831/msabookstore/bookstore_users-api/domain/users"
	"github.com/atg0831/msabookstore/bookstore_users-api/utils/crypto_utils"
	"github.com/atg0831/msabookstore/bookstore_users-api/utils/date_utils"
	"github.com/atg0831/msabookstore/bookstore_users-api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	GetAllUsers() (users.Users, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

func (s *userService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func (s *userService) GetAllUsers() (users.Users, *errors.RestErr) {
	user := &users.User{}
	return user.GetAll()
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	//현재 db에 있는 user 정보를 current에 담음
	current, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	user.Email = strings.TrimSpace(user.Email)

	//isPartial==true => PATHCH request
	//isPartial==false => PUT request
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.Status != "" {
			current.Status = user.Status
		}
	} else {
		//current에 update할 user 정보를 update
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *userService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

func (s *userService) Search(status string) (users.Users, *errors.RestErr) {
	user := &users.User{}
	//dao의 findbystatus가 user의 list를 return하고 users.Users가 []User 타입이니까
	//그대로 return하면됨
	return user.FindByStatus(status)

}

func (s *userService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	return dao, nil
}
