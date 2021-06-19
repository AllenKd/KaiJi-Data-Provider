package route

import (
	"KaiJi-Admin/internal/pkg/route/handler"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	v1.POST("/signIn", handler.UserSignIn)
}
