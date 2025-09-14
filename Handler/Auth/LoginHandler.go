package Auth

import (
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	DTOReponse "SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginPayload DTOAuth.Login

		fmt.Println("hit")

		if err := c.ShouldBind(&loginPayload); err != nil {
			fmt.Println("Bind error:", err)
			c.JSON(http.StatusBadRequest, DTOReponse.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(loginPayload); err != nil {
			fmt.Println("hit err")
			c.JSON(http.StatusBadRequest, DTOReponse.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}
		fmt.Println("hit 1")

		result, err := Auth.LoginService(loginPayload)

		fmt.Println("hit 2")

		if err != nil {
			fmt.Println("hit 3")

			formatedErr := Utils.ErrorFormater(err)
			c.JSON(formatedErr.Status, formatedErr)
			c.Abort()
			return
		}

		fmt.Println("hit 1000")

		c.JSON(http.StatusOK, DTOReponse.ResponseSuccess[DTOAuth.Auth]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    result,
		})
	}
}
