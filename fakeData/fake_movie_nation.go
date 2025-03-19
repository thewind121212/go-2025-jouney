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

func ImportFakeNationMoviePg() {

	fmt.Println("ImportFakeNation called")

	fileNation, err := os.Open("nation.json")
	if err != nil {
		fmt.Println("Error opening nation.json", err)
		panic(err)
	}

	defer fileNation.Close()

	fileNationDecoder := json.NewDecoder(fileNation)

	var nationList []_structs.Nation

	err = fileNationDecoder.Decode(&nationList)

	var idReturn string

	queryInsert := "INSERT INTO \"Nationality\" (id, name, slug, \"createdAt\", \"updatedAt\") VALUES ($1, $2, $3, $4, $5) RETURNING id"

	for _, nation := range nationList {
		id := cuid.New()
		err = repositories.Dbclient.QueryRow(queryInsert, id, nation.Name, nation.Slug, time.Now(), time.Now()).Scan(&idReturn)
		if err != nil {
			fmt.Println("Error inserting nation", err)
		} else {
			fmt.Println("Nation inserted", idReturn)
		}
	}

	fmt.Println("Write nation compoleted")

}
