package db

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	"tdaserver/pkg/utils"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// INSERT INTO games (uuid, created_at, updated_at, name, difficulty, game_state, board) VALUES (?, ?, ?, ?, ?, ?, ?)
	createGame *sql.Stmt

	// SELECT * FROM games
	getGames *sql.Stmt

	// SELECT * FROM games WHERE uuid=?
	findGame *sql.Stmt

	// UPDATE games SET updated_at=?, name=?, difficulty=?, game_state=?, board=? WHERE uuid=? RETURNING created_at
	editGame *sql.Stmt

	// DELETE FROM games WHERE uuid=?
	deleteGame *sql.Stmt
)

//go:embed initdb.sql
var initsql string

func InitDB(ctx context.Context) (*sql.DB, context.Context) {
	db, err := sql.Open("sqlite3", "file:data.db?cache=shared&mode=rwc&_journal_mode=WAL&_busy_timeout=50")
	if err != nil {
		log.Printf("Error DB open: %#v\n", err)
		return nil, context.Background()
	}

	db.SetMaxOpenConns(1)

	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Error DB ping: %#v\n", err)
		return nil, context.Background()
	}

	if len(initsql) != 0 {
		_, err := db.ExecContext(ctx, initsql)
		if err != nil {
			log.Printf("Error in DB init function: %#v\n", err)
			return nil, context.Background()
		}

		log.Printf("DB Init completed\n")
	}

	err = prepareQueries(db)
	if err != nil {
		log.Printf("Error preparing queries: %#v\n", err)
		return nil, context.Background()
	}

	end, endProgram := context.WithCancel(context.Background())

	go func() {
		<-ctx.Done()

		closeQueries()

		err := db.Close()
		log.Printf("DB Closed: %#v\n", err)

		endProgram()
	}()

	return db, end
}

func prepareQueries(db *sql.DB) (err error) {
	createGame, err = db.Prepare("INSERT INTO games (uuid, created_at, updated_at, name, difficulty, game_state, board) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	getGames, err = db.Prepare("SELECT * FROM games")
	if err != nil {
		return err
	}

	findGame, err = db.Prepare("SELECT * FROM games WHERE uuid=?")
	if err != nil {
		return err
	}

	editGame, err = db.Prepare("UPDATE games SET updated_at=?, name=?, difficulty=?, game_state=?, board=? WHERE uuid=? RETURNING created_at")
	if err != nil {
		return err
	}

	deleteGame, err = db.Prepare("DELETE FROM games WHERE uuid=?")
	if err != nil {
		return err
	}

	return nil
}

func closeQueries() {
	log.Println("Closing prepared queries")
	err := createGame.Close()
	utils.LogIfErrNotNil("Closing createGame query", err)
	err = getGames.Close()
	utils.LogIfErrNotNil("Closing getGames query", err)
	err = findGame.Close()
	utils.LogIfErrNotNil("Closing getGameByUUID query", err)
	err = editGame.Close()
	utils.LogIfErrNotNil("Closing editGame query", err)
	err = deleteGame.Close()
	utils.LogIfErrNotNil("Closing deleteGame query", err)
}
