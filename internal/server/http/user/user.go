package user

import (
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
