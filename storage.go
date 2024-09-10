package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	connStr := "postgresql://sunny3355singh:vFAOPyjRL92z@ep-royal-pine-a54hj425-pooler.us-east-2.aws.neon.tech/nikoDb?sslmode=require"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil

}

func (s *PostgressStore) init() error {
	return nil
}

func (s *PostgressStore) CreateAccountTable() error {
	query := `create table if not exists Account ( 
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_At timestamp,
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateAccount(*Account) error {
	return nil
}

func (s *PostgressStore) DeleteAccount(id int) error {
	return nil
}
func (s *PostgressStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgressStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}
