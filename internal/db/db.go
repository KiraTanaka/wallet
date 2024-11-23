package db

import (
	"database/sql"
	_ "embed"
	"fmt"

	"wallet/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type DbModels struct {
	WalletModel          *WalletDb
	WalletOperationModel *WalletOperationDb
}

var ErrorNoRows error = sql.ErrNoRows

func NewConnect(config *config.Configuration) (*DbModels, error) {
	var dbModels DbModels
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Error(err)
		return &dbModels, err
	}

	err = db.Ping()
	if err != nil {
		log.Error(err)
		return &dbModels, err
	}
	log.Info("Connection to the database is completed.")

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Error(err)
		return &dbModels, err
	}
	migrate, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Error(fmt.Sprintf("New Migrate instance error: %s", err))
	}
	err = migrate.Up()
	if err != nil {
		log.Error(fmt.Sprintf("Migration up error: %s", err))
	}

	log.Info("Verification and application of missing migrations is completed.")

	dbModels = DbModels{WalletModel: &WalletDb{Db: db},
		WalletOperationModel: &WalletOperationDb{Db: db}}

	return &dbModels, nil
}
