package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	router.DELETE("/song/:id", deleteSong)

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
	log.Debug("Post.Song")
	var request struct {
		Song  string `json:"song"`
		Group string `json:"group"`
	}
	if err := context.BindJSON(&request); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	var ng storage.NewGroup
	var ns storage.NewSong

	ng.Id = uuid.New().String()
	ng.Name = request.Group
	result := db.Table("groups").Create(ng)
	if result.Error != nil {
		log.Error(result.Error)
		log.Debug(ng)
		context.JSON(http.StatusInternalServerError, result.Error)
		return
	}
	ns.Id = uuid.New().String()
	ns.Name = request.Song
	ns.GroupId = ng.Id
	res := db.Table("songs").Create(ns)
	if res.Error != nil {
		log.Error(res.Error)
		log.Debug(ns)
		context.JSON(http.StatusInternalServerError, res.Error)
		return
	}
	context.JSON(http.StatusCreated, request)
}

func deleteSong(context *gin.Context) {
	log.Debug("Delete.Song")
	songId := context.Param("id")
	result := db.Table("songs").Where("id = ?", songId).Delete(&storage.Song{})
	if result.Error != nil {
		log.Error(result.Error)
		context.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	result = db.Table("song_details").Where("song_id = ?", songId).Delete(&storage.SongDetail{})
	if result.Error != nil {
		log.Error(result.Error)
		context.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	result = db.Table("verses").Where("song_id = ?", songId).Delete(&storage.Verse{})
	if result.Error != nil {
		log.Error(result.Error)
		context.JSON(http.StatusInternalServerError, result.Error)
		return
	}
	context.JSON(http.StatusOK, gin.H{"Deleted song with id": songId})
}
