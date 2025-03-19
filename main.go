package main

import (
	"wliafdew/go-2025/insert_genre"
	"wliafdew/go-2025/repositories"
)

func main() {
	repositories.InitDbConnection()
	insert_genre.ImportFakeGenre()
}
