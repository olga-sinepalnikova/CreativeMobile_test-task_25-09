package helpers

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetLimitAndOffset(limitStr, offsetStr string) (int, int) {
	logrus.Debug("Helpers.GetLimitAndOffset")
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	return limit, offset
}
