package Auth

import (
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SSOHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var SSORequest DTOAuth.SSO

		if err := c.ShouldBindJSON(&SSORequest); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(SSORequest); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}

		redirectUrl, err := Auth.SSOService(SSORequest)
		if err != nil {
			FormatedError := Utils.ErrorFormater(err)
			c.JSON(FormatedError.Status, FormatedError)
			c.Abort()
			return
		}

		c.JSON(http.StatusFound, DTOResponse.ResponseSuccess[string]{
			Status:  http.StatusFound,
			Message: http.StatusText(http.StatusFound),
			Data:    redirectUrl,
		})
	}
}
