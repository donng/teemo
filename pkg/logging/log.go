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

	var file *os.File
	_, err := os.Stat(logFile)
	if os.IsNotExist(err) {
		file, _ = os.Create(logFile)
	} else {
		file, _ = os.Open(logFile)
	}

	return file
}
