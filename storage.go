package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateExpanse(*Expanse) error
	DeleteExpanse(int) error
	UpdateExpanse(*Expanse) error
	GetExpanseByID(int) (*Expanse, error)
	GetExpanses() ([]*Expanse, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=expanse sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createExpanseTable()
}

func (s *PostgresStore) createExpanseTable() error {
	query := `CREATE TABLE IF NOT EXISTS expanse  (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		amount FLOAT,
		createdOn Date
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateExpanse(exp *Expanse) error {
	query := `INSERT INTO expanse(first_name, last_name, amount, createdOn)
	VALUES($1, $2, $3, $4)`

	resp, err := s.db.Query(
		query,
		exp.FirstName,
		exp.LastName,
		exp.Amount,
		exp.CreatedOn)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) UpdateExpanse(*Expanse) error {
	return nil
}

func (s *PostgresStore) DeleteExpanse(id int) error {
	return nil
}

func (s *PostgresStore) GetExpanseByID(id int) (*Expanse, error) {
	return nil, nil
}

func (s *PostgresStore) GetExpanses() ([]*Expanse, error) {
	rows, err := s.db.Query(`SELECT * FROM expanse`)
	if err != nil {
		return nil, err
	}

	expanses := []*Expanse{}
	for rows.Next() {

		expanse := new(Expanse)
		err := rows.Scan(
			&expanse.ID,
			&expanse.FirstName,
			&expanse.LastName,
			&expanse.Amount,
			&expanse.CreatedOn)

		if err != nil {
			return nil, err
		}

		expanses = append(expanses, expanse)

	}
	return expanses, nil

}
