package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

// router logic for gin
func server() (*gin.Engine, error) {
	r := gin.Default()
	r.Use(cors.Default())
	curr_path := "/home/ahmed/"
	files_data, err := ls(curr_path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, string(files_data))
	})

	r.POST("/req", func(ctx *gin.Context) {
		var data JSON_req

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(400, gin.H{"error": err})
			return
		}

		if data.Request == "cd" {
			curr_path = path.Join(curr_path, data.Path)
			files_data, err = ls(curr_path)
			if err != nil {
				ctx.JSON(400, gin.H{"error": err})
				return
			}
		}
	})

	r.GET("/download", func(ctx *gin.Context) {
		file_name := ctx.DefaultQuery("name", "")
		file_path := path.Join(curr_path, file_name)
		if _, err := os.Stat(file_path); os.IsNotExist(err) {
			ctx.JSON(404, gin.H{"error": "File not found"})
			return
		}
		ctx.FileAttachment(file_path, file_name)
	})

	r.POST("/mkdir", func(ctx *gin.Context) {
		var data JSON_req

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(400, gin.H{"error": err})
			return
		}

		folder_name := data.Path
		os.Mkdir(folder_name, 0700)
	})

	r.PUT("/", func(ctx *gin.Context) {
		ctx.String(200, string(files_data))
	})

	return r, nil
}
