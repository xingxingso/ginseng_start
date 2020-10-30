package main

import (
	"fmt"
	"ginseng_start/config"
	"ginseng_start/middlewares"
	"ginseng_start/models"
	"ginseng_start/routers"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	gin.SetMode(gin.ReleaseMode)
	fmt.Println(gin.Mode())
	if gin.Mode() != gin.DebugMode {

		gin.DisableConsoleColor()
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
		gin.DefaultErrorWriter = io.MultiWriter(f)
		// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
		log.SetOutput(gin.DefaultWriter)
	}

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})

	r := routers.SetupRouter()
	r.Use(middlewares.Cors())
	// r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//running
	r.Run()
}
