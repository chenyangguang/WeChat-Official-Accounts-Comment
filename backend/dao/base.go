package dao

/*
import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "git:Git0618@/wechat_comments?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("imit connect database fail:", err)
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(time.Hour)

	Db = db
	go func() {
		pingTime := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-pingTime.C:
				if Db != nil && Db.DB() != nil {
					Db.DB().Ping()
				}
			}
		}

	}()
}
*/
