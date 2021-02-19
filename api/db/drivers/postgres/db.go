package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	// Import of postgres driver
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/allez-chauffe/marcel/api/db/internal/db"
	"github.com/allez-chauffe/marcel/pkg/config"
)

type postgresDriver struct{}

var Driver db.Driver = new(postgresDriver)

func (driver *postgresDriver) Open() (db.Database, error) {
	pgConf := config.Default().API().DB().Postgres()

	log.Infof("Connecting to postgres database (%s:%s/%s) ...", pgConf.Host(), pgConf.Port(), pgConf.DBName())

	pg, err := sql.Open("postgres", getConnectionString(""))
	if err != nil {
		return nil, err
	}

	if err := pg.Ping(); err != nil {
		if strings.Contains(err.Error(), "database") && strings.Contains(err.Error(), "does not exist") {
			pg.Close()
			var creationError error
			if pg, creationError = createDatabase(); creationError != nil {
				return nil, fmt.Errorf("%w (failed to create it: %s", err, creationError)
			}
		} else {
			return nil, err
		}
	}

	log.Info("Postgres database connected")

	return &postgresDatabase{pg}, nil
}

type postgresDatabase struct {
	pg *sql.DB
}

func (database *postgresDatabase) Begin() (db.Transaction, error) {
	return database.pg.Begin()
}

func (database *postgresDatabase) CreateStore(newEntity func() db.Entity) (db.Store, error) {
	entity := newEntity()
	table := toSnakeCase(reflect.TypeOf(entity).Elem().Name())
	idType := reflect.TypeOf(entity.GetID()).Name()

	var postgresIDType string
	if strings.HasPrefix(idType, "int") {
		postgresIDType = "serial"
	} else if idType == "string" {
		postgresIDType = "text"
	} else {
		return nil, errors.New("Postgres database driver only supports uuid and int ID types")
	}

	log.Infof("Ensure table %s exists", table)
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS "%s" (
			"id" %s NOT NULL PRIMARY KEY,
			"data" jsonb NOT NULL
		);
	`, table, postgresIDType)

	if _, err := database.pg.Exec(query); err != nil {
		return nil, fmt.Errorf("Error while creating table '%s': %w \nQuery: %s", table, err, query)
	}

	return &postgresStore{&postgresStoreConfig{table, postgresIDType, newEntity, database.pg}, nil}, nil
}

func (database *postgresDatabase) Close() error {
	return database.pg.Close()
}

func createDatabase() (*sql.DB, error) {
	dbName := config.Default().API().DB().Postgres().DBName()
	log.Infof("Creating database %s", dbName)

	tempDB, err := sql.Open("postgres", getConnectionString("postgres"))
	if err != nil {
		return nil, err
	}
	defer tempDB.Close()

	if _, err = tempDB.Exec(fmt.Sprintf(`CREATE DATABASE "%s"`, dbName)); err != nil {
		return nil, err
	}

	pg, err := sql.Open("postgres", getConnectionString(""))
	if err != nil {
		return nil, err
	}

	return pg, pg.Ping()
}
