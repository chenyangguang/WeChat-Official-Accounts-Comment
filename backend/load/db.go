package load

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Conn *gorm.DB

func init() {
	conn, err := gorm.Open("mysql", "git:Git0618@/wechat_comments?charset=utf8&parseTime=true")
	if err != nil {
		log.Println("Init connect database fail:", err)
		return
	}
	log.Println("Load connect database work!")

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	conn.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	conn.DB().SetMaxOpenConns(100)

	// SetconnMaxLifetiment 设置连接的最大可复用时间。
	conn.DB().SetConnMaxLifetime(time.Hour)
	conn.LogMode(true)

	Conn = conn
	go func() {
		pingTime := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-pingTime.C:
				if Conn != nil && Conn.DB() != nil {
					Conn.DB().Ping()
				}
			}
		}

	}()
}
