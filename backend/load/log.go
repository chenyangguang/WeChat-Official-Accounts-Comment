package load

import (
	"fmt"
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/config"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
)

var (
	Logger *logrus.Logger
)

func InitLog() {
	goPath := os.Getenv("GOPATH")
	goPath, _ = filepath.Abs(goPath)
	logPath := filepath.Join(goPath + "/src/github.com/chenyangguang/WeChat-Official-Accounts-Comment" + "/runtime")

	err := os.MkdirAll(logPath, 0755)
	if err != nil {
		panic(err)
	}
	abLogFile := path.Join(logPath, config.LogFileName)
	fmt.Println(abLogFile)
	_, err = os.Create(abLogFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	out, err := os.OpenFile(abLogFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	Logger = logrus.New()
	Logger.Out = out
	Logger.SetLevel(logrus.DebugLevel)

	Logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
}
