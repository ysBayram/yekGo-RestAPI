package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Ogrenci struct {
	gorm.Model
	Name    string
	Surname string
	Number  int
}

type OgrenciHandler struct {
	db *gorm.DB
}

func NewOgrenciHandler(db *gorm.DB) OgrenciHandler {
	return OgrenciHandler{db}
}

func (h *OgrenciHandler) DeleteOgrenci(c *gin.Context) {
	id := c.Params.ByName("id")
	var ogrenci Ogrenci
	if err := h.db.Where("id = ?", id).First(&ogrenci).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		h.db.Delete(&ogrenci)
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	}
}
func (h *OgrenciHandler) UpdateOgrenci(c *gin.Context) {
	var ogrenci Ogrenci
	id := c.Params.ByName("id")
	if err := h.db.Where("id = ?", id).First(&ogrenci).Error; err != nil {
		c.AbortWithError(404, err)
		fmt.Println(err)
	} else {
		if err := c.BindJSON(&ogrenci); err != nil {
			c.AbortWithError(500, err)
		} else {
			h.db.Save(&ogrenci)
			c.JSON(200, ogrenci)
		}
	}
}
func (h *OgrenciHandler) CreateOgrenci(c *gin.Context) {
	var ogrenci Ogrenci
	c.BindJSON(&ogrenci)
	if err := h.db.Create(&ogrenci).Error; err != nil {
		c.AbortWithError(500, err)
		fmt.Println(err)
	} else {
		c.JSON(200, ogrenci)
	}
}
func (h *OgrenciHandler) GetOgrenciByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var ogrenci Ogrenci
	if err := h.db.Where("id = ?", id).First(&ogrenci).Error; err != nil {
		c.AbortWithError(404, err)
		fmt.Println(err)
	} else {
		c.JSON(200, ogrenci)
	}
}

func (h *OgrenciHandler) GetOgrenci(c *gin.Context) {
	var ogrenci []Ogrenci
	if err := h.db.Find(&ogrenci).Error; err != nil {
		c.AbortWithError(500, err)
		fmt.Println(err)
	} else {
		c.JSON(200, ogrenci)
	}
}
