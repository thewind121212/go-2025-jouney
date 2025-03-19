package fakeData

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/lucsky/cuid"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"wliafdew/go-2025/repositories"
	"wliafdew/go-2025/structs"
)

func getCountry() []string {

	var countryList []string

	countryFile, err := os.Open("country.json")
	if err != nil {
		panic(err)
	}

	defer countryFile.Close()

	jsonDecoder := json.NewDecoder(countryFile)
	err = jsonDecoder.Decode(&countryList)
	if err != nil {
		fmt.Println(err)
	}

	return countryList

}

func getTimeZone() []string {
	var timeZoneList []struct {
		Zone string `json:"zone"`
		Gmt  string `json:"gmt"`
		Name string `json:"name"`
	}

	fileTimezone, err := os.Open("timezone.json")
	if err != nil {
		panic(err)
	}

	defer fileTimezone.Close()

	jsonDecoder := json.NewDecoder(fileTimezone)

	err = jsonDecoder.Decode(&timeZoneList)

	if err != nil {
		fmt.Println(err)
	}

	var timeZone []string

	for _, v := range timeZoneList {
		timeZone = append(timeZone, v.Zone)
	}

	return timeZone

}

func ImportFakeUserPg() {

	countries := getCountry()
	timeZone := getTimeZone()

	//fix
	gender := [3]string{"Male", "Female", "Other"}

	file, err := os.Open("300000_users.sql")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var row []string
	var newUser []_structs.User

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Trim(line, " ") == "" {
			continue
		}
		row = append(row, line)

	}

	for _, v := range row {
		var user _structs.User
		ms := v[1 : len(v)-2]
		data := strings.Split(ms, ",")
		user.ID = cuid.New()
		user.Name = data[1]
		user.Email = data[4]
		user.Gender = data[3]
		//random time.time from age
		age, ok := strconv.Atoi(strings.Trim(data[2], " "))
		if ok != nil {
			fmt.Println(data[2])
			age = 20
		}

		//random from 0 to 2
		rn := rand.Intn(3)
		user.Gender = gender[rn]
		user.BirthDate = time.Now().AddDate(-1*age, 0, 0)
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		ranCountry := countries[rand.Intn(len(countries))]
		user.Country = ranCountry
		ranTimeZone := timeZone[rand.Intn(len(timeZone))]
		user.TimeZone = ranTimeZone
		user.Password = "$2b$10$IGmuvx8p4DqLRCrh4PYwEuGW33PgvtHeGHLW8tdMk5bFr/jzhE896"

		newUser = append(newUser, user)
	}

	//insert to db 1 first
	queryInsert := "INSERT INTO \"User\" (id, email, password, name, \"birthDate\", gender, country, \"timeZone\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"

	//using batch for each 100k

	for i, v := range newUser {
		var id string
		err = repositories.Dbclient.QueryRow(queryInsert, v.ID, v.Email, v.Password, v.Name, v.BirthDate, v.Gender, v.Country, v.TimeZone).Scan(&id)

		if err != nil {
			fmt.Println(err)
		}
		// batch 100k
		if (i+1)%100000 == 0 {
			time.Sleep(10 * time.Second)
		}
	}

}
