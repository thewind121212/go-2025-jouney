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

func fetchData(page int) {

	url := fmt.Sprintf("https://phimapi.com/v1/api/danh-sach/phim-le?limit=64&page=%d&sort_type=asc", page)

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
	totalPage := 213
	for i := 1; i <= totalPage; i++ {
		fetchData(1)
		fmt.Println(len(allMovie))
		time.Sleep(2 * time.Second)
	}

	fileMovie, err := os.Create("movie.json")

	if err != nil {
		fmt.Println("Error creating file", err)
	}

	jsonByte, err := json.Marshal(allMovie)

	fileMovie.Write(jsonByte)

}
