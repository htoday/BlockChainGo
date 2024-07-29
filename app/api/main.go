package main

import (
	"BlockChain/app/dao"
	"BlockChain/app/initialize"
	"BlockChain/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	initialize.BuildBlockChain()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"BlockChain": dao.Blockchain,
		})
	})
	r.POST("/add", func(c *gin.Context) {
		var req model.AddRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		newBlock, err := dao.GenerateBlock(dao.Blockchain[len(dao.Blockchain)-1], req.BPM)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to generate block",
			})
			return
		}
		dao.Blockchain = append(dao.Blockchain, newBlock)
		c.JSON(http.StatusOK, gin.H{
			"message": "Block added",
		})
	})
	r.Run(":8080")
}
