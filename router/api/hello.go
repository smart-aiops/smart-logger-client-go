package api

import (
	"github.com/gin-gonic/gin"
	"smart/logger/client/demo/pkg/responsex"
)

func Hello(c *gin.Context) {
	responsex.APIResponse(c, responsex.OK, "hello grpc")
}
