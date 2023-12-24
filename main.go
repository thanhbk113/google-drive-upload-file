package main

import (
	"fmt"
	"os"
	"thanhbk113/internal/config"
	"thanhbk113/util"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	util.InitFileGoogleDriverAuth()
	r := gin.Default()
	r.POST("/upload", UploadFile())
	//run server
	r.Run(":8080")
}

// UploadFile
func UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": "Error",
			})
			return
		}

		path, err := util.CreateFile(file)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": "Error",
			})
		}

		//Resize image
		// util.ResizeAndSaveImage(path, 400, 400)

		//Upload to google drive
		dir, _ := os.Getwd()
		serviceAuth := dir + config.PathGoogleDriverAuth

		url := util.UploadFileToGoogleDriver(serviceAuth, path)

		//Delete file
		err = util.DeleteFile(path)

		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": "Error",
			})
		}

		if url == "" {
			c.JSON(500, gin.H{
				"message": "Error",
			})
			return
		}

		c.JSON(200, gin.H{
			// "url":     url,
			"message": "Success",
		})
	}
}
