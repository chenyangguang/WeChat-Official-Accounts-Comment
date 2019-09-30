package middleware

import (
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/config"
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/load/log"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/silenceper/wechat"
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
		accessToken := ctx.Request.FormValue("access_token")
		if accessToken == "" {
			responseWithoutAuth(401, "Required access token", ctx)
			return
		}

		config := &wechat.Config{
			AppID:          config.AppId,
			AppSecret:      config.Secret,
			Token:          config.Token,
			EncodingAESKey: config.EncodingAESKey,
		}

		wc := wechat.NewWechat(config)
		server := wc.GetServer(ctx.Request, ctx.Writer)
        token,aErr:= server.GetAccessTokenFromServer()
        log.Logger.Info(token,aErr)

		// 传入request和responseWriter
		//server := wc.GetServer(ctx.Request, ctx.Writer)
		//log.Logger.Info("server", server.Token)
		////token, aErr := server.GetAccessToken()
		////log.Logger.Info(token, aErr)
		////设置接收消息的处理方法
		//server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//
		//	//回复消息：演示回复用户发送的消息
		//	text := message.NewText(msg.Content)
		//	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		//})
		//
		////处理消息接收以及回复
		//err := server.Serve()
		//if err != nil {
		//	log.Logger.Info("err:", err)
		//	return
		//}
		////发送回复的消息
		//server.Send()



		//token, err := context.GetAccessToken()
		//if err != nil {
		//	responseWithoutAuth(401, "Required access_token", ctx)
		//}
		//if accessToken != token {
		//	responseWithoutAuth(403, "Invalid token", ctx)
		//}
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
