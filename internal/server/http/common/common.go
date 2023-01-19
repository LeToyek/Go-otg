package common

import (
	// golang package
	"net/http"

	// external package
	"github.com/gin-gonic/gin"
)

// HandleGetCommonMessage handle get common message by given c pointer of gin.Context.
func (handler *Handler) HandleGetCommonMessage(c *gin.Context) {
	result, err := handler.common.GetCommonMessageByID(11)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}
