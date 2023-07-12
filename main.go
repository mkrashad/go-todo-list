package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model //embeded GORM struct
	name       string
	completed  bool
}



func main() {
	dsn := "host=localhost user=rashad password=1234 dbname=todo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

}
