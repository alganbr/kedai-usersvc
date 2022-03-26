package controllers

import (
	"github.com/alganbr/kedai-usersvc/internal/managers"
	"github.com/alganbr/kedai-usersvc/internal/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IUserController interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Patch(c *gin.Context)
}

type UserController struct {
	userManager managers.IUserManager
}

func NewUserController(userManager managers.IUserManager) IUserController {
	return &UserController{
		userManager: userManager,
	}
}

// Get godoc
// @Description  Get user by ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id       path      int            true  "User ID"
// @Success      200  {object}  models.User
// @Router       /user/{id} [get]
func (ctrl *UserController) Get(c *gin.Context) {
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}
	result, getErr := ctrl.userManager.Get(id)
	if getErr != nil {
		c.AbortWithStatusJSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

// Create godoc
// @Description  Create a new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request  body      models.UserRq  true  "User Request"
// @Success      201      {object}  models.User
// @Router       /user [post]
func (ctrl *UserController) Create(c *gin.Context) {
	var rq models.UserRq
	if bindErr := c.ShouldBindJSON(&rq); bindErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	result, saveErr := ctrl.userManager.Create(&rq)
	if saveErr != nil {
		c.AbortWithStatusJSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// Update godoc
// @Description  Update existing user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id       path      int            true  "User ID"
// @Param        request  body      models.UserRq  true  "User Request"
// @Success      200      {object}  models.User
// @Router       /user/{id} [put]
func (ctrl *UserController) Update(c *gin.Context) {
	var rq models.UserRq
	if bindErr := c.ShouldBindJSON(&rq); bindErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}
	result, updateErr := ctrl.userManager.Update(id, &rq)
	if updateErr != nil {
		c.AbortWithStatusJSON(updateErr.Code, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

// Patch godoc
// @Description  Update existing user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        request  body      models.UserRq  true  "User Request"
// @Success      200      {object}  models.User
// @Router       /user/{id} [patch]
func (ctrl *UserController) Patch(c *gin.Context) {
	var rq models.UserRq
	if bindErr := c.ShouldBindJSON(&rq); bindErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}
	result, patchErr := ctrl.userManager.Patch(id, &rq)
	if patchErr != nil {
		c.AbortWithStatusJSON(patchErr.Code, patchErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
