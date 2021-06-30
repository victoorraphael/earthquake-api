package main

import (
	"fmt"
	"github.com/miguelpragier/pgkebab"
	"log"
)

var (
	db *pgkebab.DBLink
)

func dbConnect() error {
	fmt.Println("connecting database")

	cn := pgkebab.ConnStringDirect("postgres://postgres:postgres123@localhost:5432/earthquake?sslmode=disable")

	const (
		connTimeout                        = 10
		execTimeout                        = 10
		connAttemptsMax                    = 10
		connAttemptsMaxMinutes             = 10
		secondsBetweenReconnectionAttempts = 10
		debugPrint                         = true
	)

	opts := pgkebab.Options(cn, connTimeout, execTimeout, connAttemptsMax, connAttemptsMaxMinutes, secondsBetweenReconnectionAttempts, debugPrint)

	_db, err := pgkebab.NewConnected(opts)

	if err != nil {
		return err
	}

	db = _db

	return nil
}

func main() {
	if err := dbConnect(); err != nil {
		log.Fatal(err)
	}

	webserviceStart()
}
