package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
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
			curr_path += data.Path + "/"
			files_data, err = ls(curr_path)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else if data.Request == "open" {
			//file_path := curr_path + data.Path
			fmt.Println("Yay")
			ctx.Header("Content-Disposition", "attachment; filename=lol.mp4")
			ctx.Header("Content-Type", "video/mp4")
			fmt.Println("done")
		}

		fmt.Println(curr_path)
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
