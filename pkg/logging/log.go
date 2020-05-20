package logging

import (
	"fmt"
	"github.com/donng/teemo/pkg/setting"
	"os"
	"time"
)

func GetLogFilePath() string {
	if setting.Setting.App.LogPath != "" {
		return setting.Setting.App.LogPath
	}
	return "runtime/logs"
}

func GetLogFileName() string {
	logPrefix := time.Now().Format(setting.Setting.App.LogTimeFormat)

	return fmt.Sprintf("%s.log", logPrefix)
}

func GetLogFile() *os.File {
	logFile := fmt.Sprintf("%s/%s", GetLogFilePath(), GetLogFileName())

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	if err != nil {
		Logger.Panicf("error open log file, err: %s", err)
	}

	return file
}
