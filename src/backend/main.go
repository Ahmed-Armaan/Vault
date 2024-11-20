package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type JSON_ls_data struct {
	Type  rune   `json:"type"`
	Title string `json:"title"`
}

type JSON_ls_res struct {
	Files     []JSON_ls_data `json:"files"`
	Curr_path string         `json:"curr_path"`
}

type JSON_req struct {
	Request string `json:"request"`
	Path    string `json:"path"`
}

func main() {
	r, err := server()
	if err != nil || r == nil {
		fmt.Println("Error :", err)
		fmt.Printf("Server could not be setup, try again")
		return
	}

	if err := r.Run(":8000"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func ls(path string) ([]byte, error) { // ls fetches files
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

	var res_data JSON_ls_res
	res_data.Files = j
	res_data.Curr_path = path

	json_data, _ := json.Marshal(res_data)
	return json_data, nil
}
