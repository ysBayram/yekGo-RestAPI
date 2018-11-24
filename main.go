package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"yekGo-RestAPI/handlers"
)

func main() {

	db, err := gorm.Open("sqlite3", "./db/gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	r := gin.Default()

	db.AutoMigrate(&handlers.Ogrenci{})

	ogrenciHandler := handlers.NewOgrenciHandler(db)

	r.GET("/ogrenci/", ogrenciHandler.GetOgrenci)
	r.GET("/ogrenci/:id", ogrenciHandler.GetOgrenciByID)
	r.POST("/ogrenci", ogrenciHandler.CreateOgrenci)
	r.PUT("/ogrenci/:id", ogrenciHandler.UpdateOgrenci)
	r.DELETE("/ogrenci/:id", ogrenciHandler.DeleteOgrenci)

	r.Run(":4141")
}
