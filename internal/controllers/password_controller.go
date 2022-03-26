package controllers

import (
	"github.com/alganbr/kedai-usersvc/internal/managers"
	"github.com/alganbr/kedai-usersvc/internal/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IPasswordController interface {
	Update(c *gin.Context)
	Validate(c *gin.Context)
}

type PasswordController struct {
	passwordManager managers.IPasswordManager
}

func NewPasswordController(passwordManager managers.IPasswordManager) IPasswordController {
	return &PasswordController{
		passwordManager: passwordManager,
	}
}

// Update godoc
// @Description  Update existing user password
// @Tags         Password
// @Accept       json
// @Produce      json
// @Param        request  body  models.UserPasswordRq  true  "User Password Request"
// @Success      200
// @Router       /password [put]
func (ctrl *PasswordController) Update(c *gin.Context) {
	var rq models.UserPasswordRq
	if bindErr := c.ShouldBindJSON(&rq); bindErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	updateErr := ctrl.passwordManager.Update(&rq)
	if updateErr != nil {
		c.AbortWithStatusJSON(updateErr.Code, updateErr)
		return
	}
	c.JSON(http.StatusOK, nil)
}

// Validate godoc
// @Description  Validate user password
// @Tags         Password
// @Accept       json
// @Produce      json
// @Param        request  body      models.ValidatePasswordRq  true  "Validate Password Request"
// @Success      200      {object}  models.User
// @Router       /password/validate [post]
func (ctrl *PasswordController) Validate(c *gin.Context) {
	var rq models.ValidatePasswordRq
	if bindErr := c.ShouldBindJSON(&rq); bindErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	user, validateErr := ctrl.passwordManager.Validate(&rq)
	if validateErr != nil {
		c.AbortWithStatusJSON(validateErr.Code, validateErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
