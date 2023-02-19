package user

import (
	"go-otg/internal/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.user.GetUserDataByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"message": "failed",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) LoginUser(c *gin.Context) {

	var user user.LoginUserModel

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.user.LoginUser(c.Request.Context(), user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": user.Username + " has been logged in",
	})
}

func (h *Handler) RegisterUser(c *gin.Context) {
	var user user.RegisterUserModel

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	err := h.user.RegisterUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": user.Username + " has been registered successfuly",
	})
}
