// Package pg implements interface Store that works with PostgeSQL.
package pg

import (
	"database/sql"
	"fmt"

	"github.com/kozyrev-m/effective-mobile-task/internal/logger"
	"go.uber.org/zap"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// Store implements interface Store and allows to interact with the PostgreSQL.
type Store struct {
	conn *sql.DB
}

// NewStore is a Store constructor.
func NewStore(dsn string) (*Store, error) {
	// create connection to db
	db, err := newPostgresDB(dsn)
	if err != nil {
		return nil, err
	}
	logger.Log.Info("connection to db established")

	s := &Store{
		conn: db,
	}
	logger.Log.Info("store successfully created")

	return s, nil
}

// Close closes store.
func (s *Store) Close() {
	if err := s.conn.Close(); err != nil {
		logger.Log.Error("db close error", zap.Error(err))
		return
	}
	logger.Log.Info("store has been closed successfully")
}

// newPostgresDB establishes a connection to the PostgreSQL database.
func newPostgresDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("db open error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}

	return db, nil
}
