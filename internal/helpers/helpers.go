package helpers

import (
	"github.com/olga-sinepalnikova/creativemobile-testtask/internal/storage"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func GetLimitAndOffset(limitStr, offsetStr string) (int, int) {
	logrus.Debug("Helpers.GetLimitAndOffset")
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	return limit, offset
}

func UpdateSongText(text, songId string, db *gorm.DB) error {
	logrus.Debug("Helpers.UpdateSongText")
	var newVerse string
	tmpCount := 1
	lines := strings.Split(text, "\n")
	var verse []string
	for _, line := range lines {
		if line == "\n" {
			newVerse = strings.Join(verse, "\n")
			res := db.Table("verses").
				Where("song_id = ? AND count = ?", songId, tmpCount).
				Update("text", newVerse)
			if res.Error != nil {
				logrus.Error(res.Error.Error())
				return res.Error
			}
			verse = nil
			tmpCount++
			continue
		}
		verse = append(verse, strings.Trim(line, " \n"))
	}
	var prevCount string
	db.Table("verses").Select("MAX(count)").Where("song_id = ?", songId).Find(&prevCount)
	if prvCnt, _ := strconv.Atoi(prevCount); prvCnt > tmpCount {
		db.Table("verses").
			Where("song_id = ? AND count > ?", songId, tmpCount).
			Delete(&storage.Verse{})
	}
	return nil
}
