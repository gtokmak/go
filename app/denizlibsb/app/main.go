package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type kavsak struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Ip   string `json:"ip"`
}

var kavsaklar = []kavsak{
	{ID: "1", Name: "DNZ-74-ULUS", Ip: "10.10.74.1"},
	{ID: "2", Name: "DNZ-75-PEKDEMİR", Ip: "10.10.75.1"},
	{ID: "3", Name: "DNZ-76-KALPMERKEZİ", Ip: "10.10.76.1"},
}

func getKavsaklar(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, kavsaklar)

}

func postKavsak(c *gin.Context) {
	var yeniKavsak kavsak
	if err := c.BindJSON(&yeniKavsak); err != nil {
		return
	}
	for _, v := range kavsaklar {
		if yeniKavsak.ID == v.ID {
			c.JSON(http.StatusBadRequest, gin.H{"mesaj": "bu id daha önce kullanılmış"})
			return
		}
	}
	kavsaklar = append(kavsaklar, yeniKavsak)
	c.IndentedJSON(http.StatusCreated, yeniKavsak)
}

func getKavsakId(c *gin.Context) {
	id := c.Param("id")

	for _, v := range kavsaklar {
		if id == v.ID {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"mesaj ": "kavsak bulunamadi."})
}

func main() {

	log.Println("Starting server...")
	router := gin.Default()
	router.GET("/kavsaklar", getKavsaklar)
	router.GET("/kavsaklar/:id", getKavsakId)
	router.POST("/kavsaklar", postKavsak)
	router.Run(":8080")

}
