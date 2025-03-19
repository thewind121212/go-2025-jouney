package fakeData

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/lucsky/cuid"
	"os"
	"time"
	"wliafdew/go-2025/repositories"
	_structs "wliafdew/go-2025/structs"
)

func ImportFakeGenrePg() {

	fmt.Println("importFakeGenre called")

	fileGenre, err := os.Open("genre.json")
	if err != nil {
		fmt.Println("Error opening genre.json", err)
		panic(err)
	}

	defer fileGenre.Close()

	fileGenreDecoder := json.NewDecoder(fileGenre)

	var nationList []_structs.Genre

	err = fileGenreDecoder.Decode(&nationList)

	var idReturn string

	queryInsert := "INSERT INTO \"MovieGenre\" (id, name, description , slug, \"createdAt\", \"updatedAt\") VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	for _, nation := range nationList {
		id := cuid.New()
		nation.Desc = "Description of " + nation.Name
		err = repositories.Dbclient.QueryRow(queryInsert, id, nation.Name, nation.Desc, nation.Slug, time.Now(), time.Now()).Scan(&idReturn)
		if err != nil {
			fmt.Println("Error inserting nation", err)
		} else {
			fmt.Println("Nation inserted", idReturn)
		}
	}

	fmt.Println("Write genre compoleted")

}
