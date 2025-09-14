package Handler

import (
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFoundRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, DTOResponse.ResponseError[*string]{
			Status:  http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		})
	}
}
