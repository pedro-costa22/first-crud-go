package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedro-costa22/first-crud-go/src/common/model"
	"github.com/pedro-costa22/first-crud-go/src/common/request"
	"github.com/pedro-costa22/first-crud-go/src/common/response"
	"github.com/pedro-costa22/first-crud-go/src/config/logger"
	"github.com/pedro-costa22/first-crud-go/src/config/rest_err"
	"github.com/pedro-costa22/first-crud-go/src/config/validation"
	"github.com/pedro-costa22/first-crud-go/src/service"
)

type AuthController struct {
	userService service.IUserService
}


func NewAuthController(service service.IUserService) *AuthController {
	return &AuthController{userService: service}
}

// Login
// @Summary Login
// @Description Authenticate user with provided email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body request.UserLoginRequest true "User login credentials"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 401 {object} rest_err.RestErr
// @Failure 404 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /login [post]
func (a *AuthController) Login(c *gin.Context) {
	var req request.UserLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
        restErr := validation.ValidateUserError(err)
		logger.Error("Error trying to validate user info", err)
        c.JSON(restErr.Code, restErr)
        return
    }

	userLogin := model.NewUserLogin(req.Email, req.Password)

	user, err := a.userService.FindByEmail(userLogin.Email)
	
	if err != nil {
		restErr := rest_err.NewNotFoundError("User or password is invalid")
		c.JSON(restErr.Code, restErr)
		return
	}

	passwordIsValid := userLogin.ValidatePassword(user.Password)
	if !passwordIsValid {
		restErr := rest_err.NewNotFoundError("User or password is invalid")
		c.JSON(restErr.Code, restErr)
		return
	}	
	
	 token, err := user.GenerateJWT()
	 if err != nil {
		restErr := rest_err.NewInternalServerError(
			fmt.Sprintf("%s", err),
		)
		c.JSON(restErr.Code, restErr)
		return
	 }

	c.Header("Authorization", token)
	response := response.LoginResponse{
		Token: token,
		User: response.LoginUserResponse{
			Name: user.Name,
			Email: user.Email,
			Age: user.Age,
			ID: user.ID.String(),
		},
	}
	c.JSON(http.StatusOK, response)
}