package crawData

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	_structs "wliafdew/go-2025/structs"
)

var allMovie []_structs.Movie

type crawData struct {
	Type     string `json:"type"`
	Page     int    `json:"page"`
	JsonFile string `json:"json_file"`
}

func fetchData(page int, movieType string) {

	url := fmt.Sprintf("https://phimapi.com/v1/api/danh-sach/%s?limit=64&page=%d&sort_type=asc", movieType, page)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responData, err := io.ReadAll(res.Body)

	var MovieItem struct {
		Data struct {
			Items []_structs.Movie `json:"items"`
		} `json:"data"`
	}

	err = json.Unmarshal(responData, &MovieItem)

	allMovie = append(allMovie, MovieItem.Data.Items...)

	if err != nil {
		fmt.Println("Error creating file", err)
	}

}

func RunCrawl() {

	movieDataCraw := []crawData{
		{Type: "phim-le", Page: 213, JsonFile: "phim-le.json"},
		{Type: "phim-bo", Page: 90, JsonFile: "phim-bo.json"},
		{Type: "tv-shows", Page: 3, JsonFile: "tv-shows.json"},
		{Type: "hoat-hinh", Page: 38, JsonFile: "hoat-hinh.json"},
	}

	fmt.Println(movieDataCraw)

	crawSelected := movieDataCraw[3]

	for i := 1; i <= crawSelected.Page; i++ {
		fetchData(i, crawSelected.Type)
		fmt.Println(len(allMovie))
		time.Sleep(2 * time.Second)
	}

	fileMovie, err := os.Create(crawSelected.JsonFile)

	if err != nil {
		fmt.Println("Error creating file", err)
	}

	jsonByte, err := json.Marshal(allMovie)

	fileMovie.Write(jsonByte)

}
