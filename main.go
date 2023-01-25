package main

import (
	"valid-password/pswdvalidation"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api")
	pswdvalidation.ValidadePassword(v1)
	r.Run()
}
