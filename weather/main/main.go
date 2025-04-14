package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Time_Inter_Temp struct {
	Time           string  `json:"time"`
	Interval       int     `json:"interval"`
	Temperature_2m float64 `json:"temperature_2m"`
}

// Structure keeps some information from api answer
type Weather struct {
	Latitude  float64         `json:"latitude"`
	Longitude float64         `json:"longitude"`
	Current   Time_Inter_Temp `json:"current"`
}

// Empty variable for weather
var W_0 Weather

// Check if the coords and datetime are in db
// false, W_0 - if not
// true, weather from db - if there are
func QueryIsinDB(coords [2]float64) (bool, Weather) {
	db, err := sql.Open("sqlite3", "weatherqueries.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	_, err = db.Exec("SELECT * FROM weatherqueries;")
	if err != nil {
		create := `CREATE TABLE weatherqueries (id integer PRIMARY KEY,
		latitude real,
		longitude real,
		time text,
		interval integer,
		temperature_2m real);`
		statement, _ := db.Prepare(create)
		statement.Exec()
	}

	todaynow := time.Now().Format("2006-01-02 15") + "%"
	todaynow_format := strings.Replace(todaynow, " ", "T", -1)
	query := `SELECT * FROM weatherqueries
		WHERE ABS(latitude-?)<1
		AND ABS(longitude-?)<1
		AND time LIKE ?;`
	rows, err := db.Query(query, coords[0], coords[1], todaynow_format)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var weathers []Weather
	for rows.Next() {
		var w Weather
		var id int
		err = rows.Scan(&id, &w.Latitude, &w.Longitude,
			&w.Current.Time, &w.Current.Interval, &w.Current.Temperature_2m)
		if err != nil {
			log.Println(err)
		}
		weathers = append(weathers, w)
	}
	if len(weathers) == 0 {
		return false, W_0
	} else {
		return true, weathers[0]
	}
}

// Insert Weather in db
func WriteinDB(w Weather) {
	db, err := sql.Open("sqlite3", "weatherqueries.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	statement, _ := db.Prepare("INSERT INTO weatherqueries (latitude, longitude, time, interval, temperature_2m) VALUES (?, ?, ?, ?, ?)")
	statement.Exec(w.Latitude, w.Longitude, w.Current.Time,
		w.Current.Interval, w.Current.Temperature_2m)
	fmt.Println("inserted in db")
}

// Returns Weather for coords from openmeteo or local db
func GiveWeather(coords [2]float64) Weather {
	if isindb, w_db := QueryIsinDB(coords); isindb {
		return w_db
	} else {
		w_API := AskAPI(coords)
		//Put weather for these coords and time into db
		WriteinDB(w_API)
		return w_API
	}
}

// Return Weather structure for coords from openmeteo
func AskAPI(coords [2]float64) Weather {
	lat := strconv.FormatFloat(coords[0], 'f', -1, 64)
	long := strconv.FormatFloat(coords[1], 'f', -1, 64)
	url := "https://api.open-meteo.com/v1/forecast?latitude=" +
		lat + "&longitude=" + long +
		"&current=temperature_2m&timezone=Europe%2FMoscow"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	w := Weather{}
	err = json.Unmarshal(body, &w)
	if err != nil {
		panic(err)
	}
	fmt.Println(w)
	return w
}

func main() {
	fmt.Println("Enter the coordinates (latitude and longitude)" +
		" to get current temperature")
	var coords [2]float64
	fmt.Scan(&coords[0], &coords[1])
	fmt.Println(GiveWeather(coords))
}
