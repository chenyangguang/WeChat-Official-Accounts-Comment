package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"log"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		random, _ := uuid.NewV4()

		ctx.Writer.Header().Set("X-Request-ID", random.String())
		ctx.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Query("access_token")

		if accessToken == "" {
			responseWithoutAuth(401, "Required grant token", ctx)
			return
		} else {
			log.Println(accessToken)
		}
		if token := ctx.Request.URL.Query().Get("token"); token == "" {
			responseWithoutAuth(402, "aaa", ctx)
			return
		}
		ctx.Next()
	}
}

func responseWithoutAuth(code int, message interface{}, ctx *gin.Context) {
	resp := map[string]interface{}{
		"code":    code,
		"message": message,
	}
	ctx.JSON(code, resp)
	ctx.Abort()
}

