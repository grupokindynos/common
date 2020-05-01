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
			c.JSON(202, gin.H{"message": "voucher daily limit exceeded", "error": err.Error(), "status": 407})
		case commonErrors.ErrorShiftMinimumAmount:
			c.JSON(411, gin.H{"message": "shift minimum amount not met", "error": err.Error(), "status": 411})
		case commonErrors.ErrorShiftDailyLimit:
			c.JSON(412, gin.H{"message": "shift daily limit exceeded", "error": err.Error(), "status": 412})
		case commonErrors.ErrorNotFound:
			c.JSON(404, gin.H{"message": "element not found in database", "error": err.Error(), "status": 1})
		case commonErrors.ErrorNoUserInformation:
			c.JSON(404, gin.H{"message": "user not found in database", "error": err.Error(), "status": 1})
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

// OpenApi
func GlobalResponse(body interface{}, c *gin.Context) *gin.Context {
	c.JSON(200, body)
	return c
}

func GlobalOpenNoAuth(c *gin.Context) *gin.Context {
	c.JSON(401, gin.H{"message": "Authentication Error", "error": errors.New("you are not authorized"), "status": 401})
	return c
}

func GlobalOpenError(result interface{}, err error, c *gin.Context) *gin.Context {
	if err != nil {
		switch err {
		case commonErrors.ErrorAssetUnavailable:
			c.JSON(400, gin.H{"message": "please try again later or contact support", "error": err.Error(), "status": 400})
		case commonErrors.ErrorServiceUnavailable:
			c.JSON(403, gin.H{"message": "the service is temporarily unavailable", "error": err.Error(), "status": 403})
		case commonErrors.ErrorFillingPaymentInformation:
			c.JSON(500, gin.H{"message": "please try again. if the problem persists contact support", "error": err.Error(), "status": 500})
		case commonErrors.ErrorObtainingRates:
			c.JSON(500, gin.H{"message": "no rates found. if the problem persists contact support", "error": err.Error(), "status": 500})
		case commonErrors.ErrorShiftNotFound:
			c.JSON(404, gin.H{"message": "please make sure you are querying the correct id", "error": err.Error(), "status": 404})
		default:
			c.JSON(500, gin.H{"message": "unknown Error. please contact support.", "error": err.Error(), "status": 500})
		}
	} else {
		c.JSON(200, gin.H{"data": result, "status": 1})
	}
	return c
}
