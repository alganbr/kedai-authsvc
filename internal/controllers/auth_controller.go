package controllers

import (
	"github.com/alganbr/kedai-authsvc/internal/managers"
	"github.com/alganbr/kedai-authsvc/internal/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IAuthController interface {
	Get(*gin.Context)
	Authenticate(*gin.Context)
}

type AuthController struct {
	authManager managers.IAuthManager
}

func NewAuthController(authManager managers.IAuthManager) IAuthController {
	return &AuthController{
		authManager: authManager,
	}
}

// Get godoc
// @Description  Get access token by ID
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        id       path      string            true  "Access Token"
// @Success      200  {object}  models.AccessToken
// @Router       /auth/{id} [get]
func (ctrl *AuthController) Get(c *gin.Context) {
	id := c.Param("id")
	result, getErr := ctrl.authManager.Get(id)
	if getErr != nil {
		c.AbortWithStatusJSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

// Authenticate godoc
// @Description  Create a new user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      models.AccessTokenRq  true  "User Request"
// @Success      201      {object}  models.AccessToken
// @Router       /auth [post]
func (ctrl *AuthController) Authenticate(c *gin.Context) {
	var rq models.AccessTokenRq
	if bindErr := c.ShouldBindJSON(&rq); bindErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	result, saveErr := ctrl.authManager.Authenticate(&rq)
	if saveErr != nil {
		c.AbortWithStatusJSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
