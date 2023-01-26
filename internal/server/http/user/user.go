package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserByID(c *gin.Context) {
	res, err := h.user.GetUserDataByID(c.Request.Context(), "tyq1")
	if err != nil {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"message": "failed",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res)

}
