package Auth

import (
	"SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/DTO/Response"
	Auth2 "SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var refreshTokenRequest Auth.RefreshTokenJson

		if err := c.ShouldBindJSON(&refreshTokenRequest); err != nil {
			errMsg := err.Error()
			c.JSON(http.StatusBadRequest, Response.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   &errMsg,
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(refreshTokenRequest); err != nil {
			c.JSON(http.StatusBadRequest, Response.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   &err,
			})
		}

		newCredential, err := Auth2.RefreshTokenService(refreshTokenRequest)

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

		c.JSON(http.StatusOK, Response.ResponseSuccess[Auth.Auth]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    &newCredential,
		})
	}
}
