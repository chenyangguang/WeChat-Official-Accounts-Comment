# 迷你公众号留言

gin 实现一个简单的 mvc 的 微信公众号留言功能后台, gorm 增\删\改\查.  MySQL存储.

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

后端完全使用 golang 实现

