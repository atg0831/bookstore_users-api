package users

import "encoding/json"

type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

//user의 list type인 Users의 경우
func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	// if isPublic {
	// 	return PublicUser{
	// 		ID:          user.ID,
	// 		DateCreated: user.DateCreated,
	// 		Status:      user.Status,
	// 	}
	// }
	// return PrivateUser{
	// 	ID: user.ID,
	// 	FirstName: user.FirstName,
	// 	...
	// }
	//위의 방법처럼 해도 되고 아래처럼 해도 된다.
	userJSON, _ := json.Marshal(user)

	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(userJSON, &publicUser)
		return publicUser
	}
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}
