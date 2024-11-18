package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

type JSON_ls_data struct {
	Type  rune   `json:"type"`
	Title string `json:"title"`
}

type JSON_cd_req struct {
	Request string `json:"request"`
	Path    string `json:"path"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	curr_path := "/home/ahmed/"
	files_data, err := ls(curr_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, string(files_data))
	})

	r.POST("/req", func(ctx *gin.Context) {
		var data JSON_cd_req

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(400, gin.H{"error": err})
			return
		}

		fmt.Println(data.Request)

		if data.Request == "cd" {
			//curr_path += data.Path + "/"
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
		fmt.Println("requested file:", file_path)
		if _, err := os.Stat(file_path); os.IsNotExist(err) {
			ctx.JSON(404, gin.H{"error": "File not found"})
			return
		}
		ctx.FileAttachment(file_path, file_name)
	})

	r.PUT("/", func(ctx *gin.Context) {
		ctx.String(200, string(files_data))
	})

	r.Run(":8000")
}

func ls(path string) ([]byte, error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var j []JSON_ls_data
	j = append(j, JSON_ls_data{
		Type:  'd',
		Title: "..",
	})

	for _, e := range dir {
		file_type := 'f'
		if e.IsDir() {
			file_type = 'd'
		}

		j = append(j, JSON_ls_data{
			Type:  file_type,
			Title: e.Name(),
		})
	}

	json_data, _ := json.Marshal(j)
	return json_data, nil
}
