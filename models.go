package main

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
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

type NullString struct {
	sql.NullString
}

type NullInt64 struct {
	sql.NullInt64
}

type Player struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	FullName NullString `db:"full_name" json:"full_name"`
	Status   NullString `json:"status"`
	Updated  NullString `json:"updated"`
	Team     NullInt64  `db:"team_id" json:"team_id"`
	Country  NullString `json:"country"`
	MMR      NullInt64  `json:"mmr"`
	Rank     NullInt64  `json:"rank"`
}

type Team struct {
	Id   int        `json:"id"`
	Name string     `json:"name"`
	Tag  NullString `json:"tag"`
}

func (nstr NullString) MarshalText() ([]byte, error) {
	if nstr.Valid {
		return []byte(nstr.String), nil
	} else {
		return nil, nil
	}
}

func (nint NullInt64) MarshalText() ([]byte, error) {
	if nint.Valid {
		return []byte(strconv.FormatInt(nint.Int64, 10)), nil
	} else {
		return nil, nil
	}
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
