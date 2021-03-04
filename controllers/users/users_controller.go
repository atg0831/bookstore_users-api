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

//url의 /:user_id param을 int로 변환해서 user의 ID 추출
func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userID, nil
}

func GetUser(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.UserService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

func GetAllUser(c *gin.Context) {
	results, err := services.UserService.GetAllUsers()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, results.Marshall(c.GetHeader("X-Public") == "true"))

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
	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateUser(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
	}

	user.ID = userID
	//Patch request 왔을 경우 isPartial = true
	isPartial := c.Request.Method == http.MethodPatch
	result, updateErr := services.UserService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	delErr := services.UserService.DeleteUser(userID)
	if delErr != nil {
		c.JSON(delErr.Status, delErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func SearchUser(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UserService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, users.Marshall((c.GetHeader("X-Public") == "true")))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
