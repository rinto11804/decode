package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateCourse(*CourseModel) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStore(cfg *Config) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", cfg.Dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	store := &PostgresStorage{
		db: db,
	}
	store, err = store.Init()
	if err != nil {
		return nil, err
	}
	log.Print(store.db.Stats())
	return store, nil
}

func (p *PostgresStorage) Init() (*PostgresStorage, error) {
	return p, nil
}

func (p *PostgresStorage) CreateCourse(e *CourseModel) error {
	_, err := p.db.Exec("INSERT INTO Event (title,description,author,userId,created_at) VALUES ($1,$2,$3)", e.Title, e.Description, e.Author, e.UserID, e.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
