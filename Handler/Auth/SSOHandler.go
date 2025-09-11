package Auth

import (
	Auth2 "SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SSOHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var SSORequest Auth2.SSOJson

		if err := c.ShouldBindJSON(&SSORequest); err != nil {
			errMsg := err.Error()
			c.JSON(http.StatusBadRequest, Response.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   &errMsg,
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(SSORequest); err != nil {
			c.JSON(http.StatusBadRequest, Response.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   &err,
			})
			c.Abort()
			return
		}

		redirectUrl, err := Auth.SSOService(SSORequest)
		if err != nil {
			errMsg := err.Error()
			c.JSON(http.StatusBadRequest, Response.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   &errMsg,
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusFound, Response.ResponseSuccess[string]{
			Status:  http.StatusFound,
			Message: http.StatusText(http.StatusFound),
			Data:    &redirectUrl,
		})
	}
}
