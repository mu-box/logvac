package authenticator

import (
	"database/sql"
	"io"

	_ "github.com/lib/pq"
)

type (
	postgresql struct {
		address string
	}
)

// NewPgDb creates a new postgres authenticator
func NewPgDb(addr string) (*postgresql, error) {
	pg := postgresql{
		address: addr,
	}
	return &pg, nil
}

func (pg postgresql) initialize() error {
	// create the tables needed to support mist authentication
	_, err := pg.exec(`
CREATE TABLE IF NOT EXISTS tokens (
	token text NOT NULL,
	PRIMARY KEY (token)
)`)
	return err
}

func (p postgresql) add(token string) error {
	_, err := p.exec("INSERT INTO tokens (token) VALUES ($1)", token)
	return err
}

func (p postgresql) remove(token string) error {
	_, err := p.exec("DELETE FROM tokens WHERE token = $1", token)
	return err
}

func (p postgresql) valid(token string) bool {
	r, err := p.query("select * FROM tokens WHERE token = $1", token)
	if err != nil {
		return false
	}
	// if there are any results then we are valid
	return r.Next()
}

func (p postgresql) connect() (*sql.DB, error) {
	return sql.Open("postgres", string(p.address))
}

// this could really be optimized a lot. instead of opening a new
// conenction for each query, it should reuse connections
func (p postgresql) query(query string, args ...interface{}) (*sql.Rows, error) {
	client, err := p.connect()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return client.Query(query, args...)
}

// This could also be optimized a lot
func (p postgresql) exec(query string, args ...interface{}) (sql.Result, error) {
	client, err := p.connect()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return client.Exec(query, args...)
}

func (p postgresql) exportLogvac(exportWriter io.Writer) error {
	return nil
}

func (p postgresql) importLogvac(importReader io.Reader) error {
	return nil
}
