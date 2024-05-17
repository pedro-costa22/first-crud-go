package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedro-costa22/first-crud-go/src/common/request"
	"github.com/pedro-costa22/first-crud-go/src/config/logger"
	"github.com/pedro-costa22/first-crud-go/src/config/rest_err"
	"github.com/pedro-costa22/first-crud-go/src/config/validation"
	"github.com/pedro-costa22/first-crud-go/src/service"
	"github.com/pedro-costa22/first-crud-go/src/view"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(service service.IUserService) *UserController {
	return &UserController{userService: service}
}

func (u *UserController) Create(c *gin.Context) {
	var req request.UserCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
        restErr := validation.ValidateUserError(err)
		logger.Error("Error trying to validate user info", err)
        c.JSON(restErr.Code, restErr)
        return
    }

	if _, err := u.userService.FindByEmail(req.Email); err == nil {
		restErr := rest_err.NewRestErr(
            "email already exists",
			"409",
			http.StatusConflict,
			nil,
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

	c.JSON(http.StatusCreated, view.CreateUserView(*user));
}

func (u *UserController) FindByID(c *gin.Context) {
	id := c.Param("id")
	user, err := u.userService.FindByID(id)

	if err != nil {
		restErr := rest_err.NewNotFoundError("User not found")
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, view.CreateUserView(user))
}

func (u *UserController) FindByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := u.userService.FindByEmail(email)
	
	if err != nil {
		restErr := rest_err.NewNotFoundError("User not found")
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, view.CreateUserView(user))
}

func (u *UserController) Update(c *gin.Context) {
    var updates map[string]interface{}
    if err := c.ShouldBindJSON(&updates); err != nil {
        restErr := validation.ValidateUserError(err)
        logger.Error("Error trying to validate user info", err)
        c.JSON(http.StatusBadRequest, restErr)
        return
    }

    id := c.Param("id")
    user, err := u.userService.Update(id, updates)
    if err != nil {
        restErr := rest_err.NewInternalServerError(err.Error())
        c.JSON(restErr.Code, restErr)
        return
    }

    c.JSON(http.StatusOK, view.CreateUserView(user))
}

func (u *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := u.userService.Delete(id)

	if err != nil {
		restErr := rest_err.NewNotFoundError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusNoContent)
}