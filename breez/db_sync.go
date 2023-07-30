package breez

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/sync/errgroup"
	"os"
)

const (
	// TODO: create on first run
	createStorage = "CREATE TABLE IF NOT EXISTS storage (file TEXT PRIMARY KEY, content BYTEA);"

	selectContents = "SELECT file, content FROM storage;"
	upsertContent  = "INSERT INTO storage (file, content) VALUES ($1, $2) ON CONFLICT (file) DO UPDATE SET content = $2;"
)

var (
	sqliteFiles = []string{
		"storage.sql",
		"sync_storage.sql",
	}
)

type dbSync struct {
	postgres *sql.DB
}

func initDbSync(con string) (*dbSync, error) {
	if con == "" {
		return &dbSync{}, nil
	}

	db, err := sql.Open("pgx", con)
	if err != nil {
		return nil, err
	}

	return &dbSync{
		postgres: db,
	}, nil
}

func (db *dbSync) close() error {
	if db.postgres == nil {
		return nil
	}

	return db.postgres.Close()
}

func (db *dbSync) download() error {
	if db.postgres == nil {
		return nil
	}

	// Files exist
	if _, err := os.Stat(sqliteFiles[0]); err == nil {
		return nil
	}

	rows, err := db.postgres.Query(selectContents)
	if err != nil {
		return err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var file string
		var content []byte

		if err = rows.Scan(&file, &content); err != nil {
			return err
		}
		if err = os.WriteFile(file, content, 0600); err != nil {
			return err
		}
	}

	return nil
}

func (db *dbSync) upload() error {
	if db.postgres == nil {
		return nil
	}

	tx, err := db.postgres.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	eg, _ := errgroup.WithContext(context.Background())

	for _, file := range sqliteFiles {
		file := file
		eg.Go(func() error {
			return db.uploadSqlite(tx, file)
		})
	}

	err = eg.Wait()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (db *dbSync) uploadSqlite(tx *sql.Tx, file string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	_, err = tx.Exec(upsertContent, file, content)

	return err
}
