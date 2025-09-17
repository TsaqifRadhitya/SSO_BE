package Auth

import (
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyAccessHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var VerifyAccessRequest DTOAuth.VerifyAccess

		if err := c.ShouldBind(&VerifyAccessRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(VerifyAccessRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}
		accesGranted, applicationName := Auth.VerifyAccessService(VerifyAccessRequest)
		if !accesGranted && applicationName == "" {
			c.JSON(http.StatusNotFound, DTOResponse.ResponseError[*interface{}]{
				Status:  http.StatusNotFound,
				Message: http.StatusText(http.StatusNotFound),
			})
			c.Abort()
			return
		}

		type verifyAccessData struct {
			ApplicationName string `json:"application_name"`
		}

		if !accesGranted && applicationName != "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, DTOResponse.ResponseSuccess[verifyAccessData]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data: verifyAccessData{
				ApplicationName: applicationName,
			},
		})
	}
}
