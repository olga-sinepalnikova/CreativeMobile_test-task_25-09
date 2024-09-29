package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olga-sinepalnikova/creativemobile-testtask/config"
	"github.com/olga-sinepalnikova/creativemobile-testtask/internal/helpers"
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
	router := gin.Default()

	router.GET("/lib", getLib)
	router.GET("/song/:id", getSong)
	router.POST("/song/", postSong)

	err = router.Run()
	if err != nil {
		log.Error(err)
	}

}

func getLib(context *gin.Context) {
	log.Debug("Get.Lib")
	limit, offset := helpers.GetLimitAndOffset(context.Query("limit"), context.Query("offset"))

	result := db.Table("songs").
		Select("songs.name, sd.release_date, sd.link, v.text," +
			" (SELECT g.name FROM groups g WHERE songs.group_id=g.id)").
		Joins("INNER JOIN song_details sd ON sd.song_id = songs.id").
		Joins("INNER JOIN verses v ON v.song_id = sd.song_id AND count=1")

	if limit != 0 {
		result = result.Limit(limit)
	}
	if offset != 0 {
		result = result.Offset(offset)
	}
	sqlResult, err := result.Rows()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	var response []storage.SongResponse
	for sqlResult.Next() {
		var tmp storage.SongResponse
		err = sqlResult.Scan(&tmp.Name, &tmp.ReleaseDate, &tmp.Link, &tmp.Text, &tmp.Group)
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

func getSong(context *gin.Context) {
	log.Debug("Get.Song")
	limit, offset := helpers.GetLimitAndOffset(context.Query("limit"), context.Query("offset"))
	result := db.Table("verses v").
		Select("v.count, v.text").
		Joins("INNER JOIN songs s ON s.id = v.song_id").
		Where("s.id = ?", context.Param("id"))
	if limit != 0 {
		result = result.Limit(limit)
	}
	if offset != 0 {
		result = result.Offset(offset)
	}
	sqlResult, err := result.Rows()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	var response []storage.SongTextResponse
	for sqlResult.Next() {
		var tmp storage.SongTextResponse
		err = sqlResult.Scan(&tmp.Count, &tmp.Text)
		if err != nil {
			log.Error(err)
			context.JSON(http.StatusInternalServerError, err)
			return
		}
		response = append(response, tmp)
	}
	context.JSON(http.StatusOK, response)
}

func postSong(context *gin.Context) {

}
