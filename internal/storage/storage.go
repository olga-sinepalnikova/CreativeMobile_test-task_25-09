package storage

import (
	"github.com/google/uuid"
	"github.com/olga-sinepalnikova/creativemobile-testtask/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Database struct {
	Song
	Group
	SongDetail
	Verse
}

type Song struct {
	gorm.Model
	ID      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name    string
	GroupID uuid.UUID `gorm:"type:uuid"`
}

type Group struct {
	gorm.Model
	ID   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	name string
}

type SongDetail struct {
	gorm.Model
	songId      uuid.UUID `gorm:"type:uuid"`
	releaseDate time.Time
	link        string
}

type Verse struct {
	gorm.Model
	songId uuid.UUID `gorm:"type:uuid"`
	id     int
	text   string
}

type SongResponse struct {
	Name        string
	Group       string
	ReleaseDate string
	Link        string
	Text        string
}

type SongRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type NewGroup struct {
	Id   string
	Name string
}

type NewSong struct {
	Id      string
	Name    string
	GroupId string
}

type SongTextResponse struct {
	Count int
	Text  string
}

func New(c config.Config) (*gorm.DB, error) {
	logrus.Debug("Storage.New")
	db, err := gorm.Open(postgres.Open(c.PostgresConn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	logrus.Info("Connected to postgres")
	err = db.AutoMigrate(&Song{})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	err = db.AutoMigrate(&SongDetail{})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	err = db.AutoMigrate(&Group{})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	err = db.AutoMigrate(&Verse{})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return db, nil
}
