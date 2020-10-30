//package main
//
//import (
//	"gorm.io/driver/mysql"
//	"github.com/gin-gonic/gin"
//	"gorm.io/gorm"
//	_ "ginseng/routers"
//)
//
//func main() {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/api?charset=utf8mb4&parseTime=True&loc=Local"
//
//	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	r := gin.Default()
//	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//}