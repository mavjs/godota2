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
	Id       int
	Name     string
	FullName NullString `db,json:"full_name"`
	Status   NullString `db,json:"status"`
	Updated  NullString
	Team     NullInt64 `db,json:"team_id"`
	Country  NullString
	MMR      NullInt64
	Rank     NullInt64
}

type Team struct {
	Id   int
	Name string
	Tag  NullString
}

func (nstr NullString) MarshalText() ([]byte, error) {
	if nstr.Valid {
		return []byte(nstr.String), nil
	} else {
		return []byte("null"), nil
	}
}

func (nint NullInt64) MarshalText() ([]byte, error) {
	if nint.Valid {
		return []byte(strconv.FormatInt(nint.Int64, 10)), nil
	} else {
		return []byte("null"), nil
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
