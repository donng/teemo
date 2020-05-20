package logging

import (
	"fmt"
	"github.com/donng/teemo/pkg/setting"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"os"
	"time"
)

func GetAccessWriter() *rotatelogs.RotateLogs {
	logf, err := rotatelogs.New("runtime/logs/access_%Y%m%d.log")
	if err != nil {
		panic(err)
	}
	return logf
}

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
	file, err := os.OpenFile("runtime/logs/access.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	if err != nil {
		Logger.Panicf("error open log file, err: %s", err)
	}

	return file
}
