package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olga-sinepalnikova/creativemobile-testtask/config"
	"github.com/olga-sinepalnikova/creativemobile-testtask/internal/handler"
	"github.com/olga-sinepalnikova/creativemobile-testtask/internal/storage"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

func main() {
	log.Info("Configuration in process...")
	conf, err := config.New()
	if err != nil {
		log.Error(err)
	}

	log.Info("Connecting to database...")
	db, err = storage.New(*conf)
	if err != nil {
		log.Error(err)
	}
	//var songs []storage.Song
	//var songDetail []storage.SongDetail
	//var groups []storage.Group
	//var verses []storage.Verse

	log.Info("Starting server...")
	router := handler.New()

	router.GET("/lib", getLib)

	err = router.Run()
	if err != nil {
		log.Error(err)
	}

}

func getLib(context *gin.Context) {
	log.Debug("Get.Lib")
	result, err := db.Table("songs").
		Select("songs.name, sd.release_date, sd.link, v.text," +
			" (SELECT g.name FROM groups g WHERE songs.group_id=g.id)").
		Joins("INNER JOIN song_details sd ON sd.song_id = songs.id").
		Joins("INNER JOIN verses v ON v.song_id = sd.song_id AND count=1").Rows()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	if err := result.Err(); err != nil {
		context.JSON(http.StatusNotFound, "Nothing found")
		return
	}

	var response []storage.SongResponse
	for result.Next() {
		var tmp storage.SongResponse
		err = result.Scan(&tmp.Name, &tmp.ReleaseDate, &tmp.Link, &tmp.Text, &tmp.Group)
		if err != nil {
			log.Error(err)
			context.JSON(http.StatusInternalServerError, err)
			return
		}
		tmp.ReleaseDate = tmp.ReleaseDate[0:10]
		response = append(response, tmp)
		log.Debug("Appended: ", tmp)
	}
	context.JSON(http.StatusOK, response)
}
