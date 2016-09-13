package main

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return
}

type Player struct {
	Id       int
	Name     string
	FullName string         `db,json:"full_name"`
	Status   sql.NullString `db,json:"status"`
	Updated  sql.NullString
	Team     sql.NullInt64 `db,json:"team_id"`
	Country  sql.NullString
	MMR      sql.NullInt64
	Rank     sql.NullInt64
}

type Team struct {
	Id   int
	Name string
	Tag  sql.NullString
}

func Players() (players []Player, err error) {
	rows, err := Db.Queryx("SELECT id, name, full_name, status, updated, team_id, country, mmr, rank FROM rosters_player")
	if err != nil {
		return
	}
	for rows.Next() {
		player := Player{}
		if err = rows.StructScan(&player); err != nil {
			return
		}
		players = append(players, player)
	}
	rows.Close()
	return
}
