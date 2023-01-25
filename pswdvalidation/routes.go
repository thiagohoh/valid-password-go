package pswdvalidation

import (
	"github.com/gin-gonic/gin"
)

func ValidadePassword(router *gin.RouterGroup) {
	router.POST("/verify", PasswordValidation)
}
