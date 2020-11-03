package main

import (
	"fmt"
	"ginseng_start/config"
	"ginseng_start/middlewares"
	"ginseng_start/models"
	"ginseng_start/routers"
	"io"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

var err error

func InitializeLogging(mode string) {
	var f io.Writer
	var level logrus.Level
	var formatter logrus.Formatter
	if mode != gin.ReleaseMode {
		f = os.Stdout
		level = logrus.DebugLevel
		formatter = &logrus.TextFormatter{
			FullTimestamp: true,
		}
	} else {
		f, err = os.OpenFile("ginseng.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
			f = os.Stdout
		}
		// todo:
		// level = logrus.WarnLevel
		level = logrus.DebugLevel
		formatter = &logrus.JSONFormatter{
			// DataKey:     "context",
			// PrettyPrint: true,
		}
	}

	// // don't Close here
	// defer f.Close()

	if formatter, ok := formatter.(*logrus.TextFormatter); ok {
		formatter.TimestampFormat = "2006-01-02 15:04:05"
	}
	if formatter, ok := formatter.(*logrus.JSONFormatter); ok {
		formatter.TimestampFormat = "2006-01-02 15:04:05"
	}

	logrus.SetOutput(f)
	logrus.SetLevel(level)
	logrus.SetFormatter(formatter)
}

func main() {
	// gin.SetMode(gin.ReleaseMode)

	InitializeLogging(gin.Mode())

	// fmt.Println(gin.Mode())
	// if gin.Mode() != gin.DebugMode {
	//
	// 	gin.DisableConsoleColor()
	// 	f, _ := os.Create("gin.log")
	// 	gin.DefaultWriter = io.MultiWriter(f)
	// 	gin.DefaultErrorWriter = io.MultiWriter(f)
	// 	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 	log.SetOutput(gin.DefaultWriter)
	// }

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})

	r := gin.New()
	// r := routers.SetupRouter()
	r.Use(middlewares.Cors())
	// r.Use(gin.Logger())
	// rLog := logrus.New()
	r.Use(middlewares.Logger(), gin.Recovery())
	// r.Use(gin.Recovery())

	middlewares.NewJwt(r)
	r.Use(middlewares.Jwt())

	routers.SetupRouter(r)
	//running
	r.Run()
}
