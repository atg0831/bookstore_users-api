package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/atg0831/msabookstore/bookstore_users-api/domain/users"
	"github.com/atg0831/msabookstore/bookstore_users-api/services"
	"github.com/atg0831/msabookstore/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(user)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	//c.ShouldBindJSON이 위의 코드들을 함축한 것이다.
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(err.Error())
		return
	}
	newUser, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, newUser)

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")

}
