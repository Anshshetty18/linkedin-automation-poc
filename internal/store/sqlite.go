package store

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

func New(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	s := &Store{db: db}
	return s, s.init()
}

func (s *Store) init() error {
	schema := `
	CREATE TABLE IF NOT EXISTS connections (
		profile_url TEXT UNIQUE,
		status TEXT,
		created_at TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS messages (
		profile_url TEXT UNIQUE,
		message TEXT,
		created_at TIMESTAMP
	);`
	_, err := s.db.Exec(schema)
	return err
}

func (s *Store) HasConnection(url string) (bool, error) {
	var c int
	err := s.db.QueryRow("SELECT COUNT(1) FROM connections WHERE profile_url = ?", url).Scan(&c)
	return c > 0, err
}

func (s *Store) SaveConnection(url, status string) error {
	_, err := s.db.Exec("INSERT OR IGNORE INTO connections VALUES (?, ?, ?)", url, status, time.Now())
	return err
}

func (s *Store) HasMessage(url string) (bool, error) {
	var c int
	err := s.db.QueryRow("SELECT COUNT(1) FROM messages WHERE profile_url = ?", url).Scan(&c)
	return c > 0, err
}

func (s *Store) SaveMessage(url, msg string) error {
	_, err := s.db.Exec("INSERT OR IGNORE INTO messages VALUES (?, ?, ?)", url, msg, time.Now())
	return err
}
