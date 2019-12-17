package responses

import (
	"errors"
	"github.com/gin-gonic/gin"
	commonErrors "github.com/grupokindynos/common/errors"
)

// GlobalResponseError is used to wrap all the errored API responses under the same model.
// Automatically detect if there is an error and return status and code according
func GlobalResponseError(result interface{}, err error, c *gin.Context) *gin.Context {
	if err != nil {
		switch err {
		case commonErrors.ErrorVoucherLimit:
			c.JSON(202, gin.H{"message": "voucher daily limit exceeded", "error": err.Error(), "status":407})
		default:
			c.JSON(500, gin.H{"message": "Error", "error": err.Error(), "status": -1})
		}
	} else {
		c.JSON(200, gin.H{"data": result, "status": 1})
	}
	return c
}

// GlobalResponseNoAuth is used to wrap all non-auth API responses under the same model.
func GlobalResponseNoAuth(c *gin.Context) *gin.Context {
	c.JSON(401, gin.H{"message": "Error", "error": errors.New("you are not authorized"), "status": -1})
	return c
}

// GlobalResponseMRT is used to wrap all responses using a MRT token
func GlobalResponseMRT(header string, body string, c *gin.Context) *gin.Context {
	c.Header("service", header)
	c.JSON(200, body)
	return c
}
