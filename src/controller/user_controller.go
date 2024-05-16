package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedro-costa22/first-crud-go/src/common/interfaces"
	"github.com/pedro-costa22/first-crud-go/src/common/request"
	"github.com/pedro-costa22/first-crud-go/src/config/rest_err"
)

type UserController struct {
	userService interfaces.IUserService
}

func NewUserController(service interfaces.IUserService) *UserController {
	return &UserController{userService: service}
}

func (u *UserController) Create(c *gin.Context) {
	var req request.UserCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
        restErr := rest_err.NewBadRequestError(
            fmt.Sprintf("There are some incorrect fields: %s", err.Error()),
        )
        c.JSON(restErr.Code, restErr)
        return
    }

	user, err := u.userService.Create(req)
	if err != nil {
		restErr := rest_err.NewInternalServerError(
            fmt.Sprintf("Failed to create user: %s", err.Error()),
        )
        c.JSON(restErr.Code, restErr)
        return
	}

	c.JSON(http.StatusCreated, user)
}

func (u *UserController) FindByID(c *gin.Context) {
	fmt.Println("passou aqui")
}

func (u *UserController) FindByEmail(c *gin.Context) {
	fmt.Println("passou aqui")
}

func (u *UserController) Update(c *gin.Context) {
	fmt.Println("passou aqui")
}

func (u *UserController) Delete(c *gin.Context) {
	fmt.Println("passou aqui")
}