package handler

import (
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/config"
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/load"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

// AccessTokenResp access_token 返回时的结构
// 微信开发文档 https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
// {"access_token":"ACCESS_TOKEN","expires_in":7200}
// {"errcode":40013,"errmsg":"invalid appid"}

// AccessTokenRespErr 获取access_token时失败返回结构
/*
-1	系统繁忙，此时请开发者稍候再试
40001	AppSecret错误或者AppSecret不属于这个公众号，请开发者确认AppSecret的正确性
40002	请确保grant_type字段值为client_credential
40164	调用接口的IP地址不在白名单中，请在接口IP白名单中进行设置。（小程序及小游戏调用不要求IP地址在白名单内。）
*/
/*
0	请求成功
*/
func SendMsg(c *gin.Context) {
	//配置微信参数
	config := &wechat.Config{
		AppID:          config.AppId,
		AppSecret:      config.Secret,
		Token:          config.Token,
		EncodingAESKey: config.EncodingAESKey,
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		load.Logger.Info(err)
		return
	}
	//发送回复的消息
	server.Send()
}
