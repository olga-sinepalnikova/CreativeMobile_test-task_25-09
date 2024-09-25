package storage

import (
	"github.com/olga-sinepalnikova/creativemobile-testtask/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(c config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.PostgresConn), &gorm.Config{})
	if err != nil {
		// todo: use logrus
		return nil, err
	}
	// todo: use logrus
	return db, nil
}
