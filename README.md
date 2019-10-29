# 迷你公众号留言
[![Go Report Card](https://travis-ci.org/silenceper/wechat.svg?branch=master)](https://travis-ci.org/chenyangguang/WeChat-Official-Accounts-Comment)
[![Go Report Card](https://goreportcard.com/badge/github.com/chenyangguang/WeChat-Official-Accounts-Comment)](https://goreportcard.com/report/github.com/chenyangguang/WeChat-Official-Accounts-Comment)


gin 实现一个简单的 mvc 的 微信公众号留言功能后台, gorm 增\删\改\查.  MySQL存储.

## Require 
```
1. go get github.com/gin-gonic/gin
2. go get github.com/gomodule/redigo/redis
3. go get github.com/bradfitz/gomemcache/memcache
4. go get golang.org/x/crypto/pkcs12c
5. go get github.com/satori/go.uuid
6. go get github.com/silenceper/wechat
```


## 目录结构 

```
├── LICENSE
├── README.md
├── backend
│   ├── comment
│   ├── config
│   │   ├── cache.go
│   │   ├── global.go
│   │   └── weixin.go
│   ├── controller
│   │   ├── base.go
│   │   └── comment.go
│   ├── dao
│   │   ├── comment.go
│   │   └── user.go
│   ├── gin.log
│   ├── handler
│   │   └── wechat.go
│   ├── load
│   │   ├── cache.go
│   │   ├── db.go
│   │   └── log
│   ├── main.go
│   ├── middleware
│   │   ├── auth.go
│   │   └── log.go
│   ├── router
│   │   └── router.go
│   └── vendor
│       └── vendor.json
├── database
│   └── comments_db.sql
├── frontend
│   ├── app.js
│   ├── app.json
│   ├── app.wxss
│   ├── images
│   │   ├── 2.jpg
│   │   ├── zan1.png
│   │   ├── zan2.png
│   │   └── 主页.png
│   ├── pages
│   │   ├── artical
│   │   ├── index
│   │   ├── logs
│   │   ├── lookmessage
│   │   ├── message
│   │   ├── myartical
│   │   ├── mycenter
│   │   ├── select
│   │   └── write
│   ├── project.config.json
│   ├── sitemap.json
│   ├── utils
│   │   └── util.js
│   ├── weui.wxss
│   └── wxSearch
│       ├── images
│       ├── wxSearch.js
│       ├── wxSearch.wxml
│       └── wxSearch.wxss
└── runtime
    └── gin.log
```
## Install

```
git clone https://github.com/chenyangguang/WeChat-Official-Accounts-Comment.git

```

或者

```
go get github.com/chenyangguang/WeChat-Official-Accounts-Comment
```

##  Init Database 初始化数据库表

创建留言数据库，建立留言相关的表:  database/comments_db.sql


## 运行

linux

```
cd backend && go run main.go

```
or


```
go build -o comment main.go && ./comment
```


windows
```
go build -o comment main.go && ./comment.exe

```

## 实现

1. 后端完全使用 golang 实现
2. 前端绝大部分是借鉴[luxiangqiang ](https://github.com/luxiangqiang/WeiXin_MessageApplet "luxiangqiang") 的界面

