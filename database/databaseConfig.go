package database

import (
	"apiCadastro/config"
	"apiCadastro/server"
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb" //bblablalba
)

// NewSQL ...
func NewSQL(conf *config.SQL) (*SQLStr, error) {
	s := &SQLStr{}
	return s, s.UpdateConfig(conf)
}

// UpdateConfig ...
func (s *SQLStr) UpdateConfig(conf *config.SQL) error {
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			return err
		}
	}
	s.conf = conf

	return s.Connect()
}

// Connect ...
func (s *SQLStr) Connect() error {
	var err error
	s.db, err = sql.Open("sqlserver", s.url.String())
	if err != nil {
		return err
	}
	return s.db.Ping()
}

// ListenAndServe ...
func (d *ConectionSqlConfig) ListenAndServe() error {
	return d.server.ListenAndServe()
}

// NewConectionSql ...
func NewConectionSql(filePath string) (*ConectionSqlConfig, error) {
	c := &ConectionSqlConfig{}

	confChan := make(chan config.Config)

	firstRead := make(chan error)
	go func() {
		conf := <-confChan
		var err error
		l, err := NewSQL(&conf.Linx)
		if err != nil {
			firstRead <- err
			return
		}
		c.linx = (*linx)(l)
		if err := (*SQLStr)(c.linx).Connect(); err != nil {
			firstRead <- err
			return
		}

		c.server = server.NewServer(conf, c.Handlers())

		firstRead <- nil

		for conf := range confChan {
			if err := (*SQLStr)(c.linx).UpdateConfig(&conf.Linx); err != nil {
				log.Fatal(err)
			}

			c.server = server.NewServer(conf, c.Handlers())
		}
	}()

	_, err := config.LoadYaml(filePath, func(conf config.Config) {
		confChan <- conf
	})
	if err != nil {
		return nil, err
	}
	if err := <-firstRead; err != nil {
		return nil, err
	}

	return c, nil
}
