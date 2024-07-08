package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Salikhov079/military-ai/config"
	u "github.com/Salikhov079/military-ai/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db    *sql.DB
	Ais u.Ai
}

func NewPostgresStorage() (u.InitRoot, error) {
	config := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword,
		config.PostgresHost, config.PostgresPort,
		config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{Db: db}, err
}


func (s *Storage) Ai() u.Ai {
	if s.Ais == nil {
		s.Ais = &AiStorage{s.Db}
	}

	return s.Ais
}
