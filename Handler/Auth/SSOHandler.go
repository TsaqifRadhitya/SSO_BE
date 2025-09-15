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
		userId, _ := c.Get("User")
		SSORequest := DTOAuth.SSO{
			UserId: userId.(string),
		}

		if err := c.ShouldBind(&SSORequest); err != nil {
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

		c.JSON(http.StatusOK, DTOResponse.ResponseSuccess[string]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    redirectUrl,
		})
	}
}
