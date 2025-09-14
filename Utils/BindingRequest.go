package Utils

import "github.com/gin-gonic/gin"

func BindAll(c *gin.Context, dto interface{}) error {
	if err := c.ShouldBind(dto); err != nil {
		return err
	}
	return nil
}

func Bind(c *gin.Context, dto interface{}, option ...string) error {
	for _, opt := range option {
		switch opt {
		case "body":
			if err := c.ShouldBind(dto); err != nil {
				return err
			}
		case "header":
			if err := c.ShouldBindHeader(dto); err != nil {
				return err
			}
		case "query":
			if err := c.ShouldBindQuery(dto); err != nil {
				return err
			}
		case "param":
			if err := c.ShouldBindUri(dto); err != nil {
				return err
			}
		}
	}
	return nil
}
