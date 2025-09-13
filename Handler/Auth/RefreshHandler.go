package Auth

import (
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	Auth2 "SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var refreshTokenRequest DTOAuth.RefreshToken

		if err := c.ShouldBindJSON(&refreshTokenRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(refreshTokenRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
		}

		newCredential, err := Auth2.RefreshTokenService(refreshTokenRequest)

		if err != nil {
			formatedError := Utils.ErrorFormater(err)
			c.JSON(formatedError.Status, formatedError)
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, DTOResponse.ResponseSuccess[DTOAuth.Auth]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    newCredential,
		})
	}
}
