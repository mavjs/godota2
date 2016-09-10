package main

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return
}

type Player struct {
	Id       int
	Name     string
	FullName string `db,json:"full_name"`
	Status   string `db,json:"status"`
	Updated  time.Time
	Team     int `db,json:"team_id"`
	Country  string
	MMR      int
	Rank     int
}

type Team struct {
	Id   int
	Name string
	Tag  string
}

func Players() (player []Player, err error) {
	rows, err := Db.Query("SELECT id, name, full_name, status, updated, team_id, country, mmr, rank FROM rosters_player")
	if err != nil {
		return
	}
	for rows.Next() {
		player := Player{}
		if err = rows.Scan(&player.Id, &player.Name, &player.FullName, &player.Status, &player.Updated, &player.Team, &player.Country, &player.MMR, &player.Rank); err != nil {
			return
		}
		players = append(players, player)
	}
	rows.Close()
	return
}
