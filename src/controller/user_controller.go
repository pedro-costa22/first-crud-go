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

// Create a new user
// @Summary Create a new user
// @Description Create a new user with the provided user information
// @Tags Users
// @Accept json
// @Produce json
// @param request body  request.UserCreateRequest true "User Data for create"
// @Success 200 {object} response.UserCreateResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /users [post]
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

// Find a user by ID
// @Summary Find a user by ID
// @Description Retrieve a user's information by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Security Bearer
// @Success 200 {object} response.UserCreateResponse
// @Failure 404 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /users/{id} [get]
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

// Find a user by Email
// @Summary Find a user by Email
// @Description Retrieve a user's information by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Security Bearer
// @Success 200 {object} response.UserCreateResponse
// @Failure 404 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /getUserByEmail/{email} [get] 
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

// User Updated 
// @Summary User Updated
// @Description Retrieve a user's information by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @param request body  request.UserUpdateRequest true "User Data for update"
// @Security Bearer
// @Success 200 {object} response.UserCreateResponse
// @Failure 404 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /users/{id} [patch]
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

// User delete 
// @Summary User delete
// @Description Retrieve a user's information by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Security Bearer
// @Success 200 {int} http.StatusNoContent
// @Failure 404 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /users/{id} [delete]
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