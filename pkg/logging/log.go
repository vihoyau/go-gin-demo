package logging

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-demo/pkg/file"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

func Setup() {
	var err error
	// 获取 基础路径与文件名
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
}
func Info(v ...interface{}) {
	setPrefix(INFO)
}
func Warning(v ...interface{}) {
	setPrefix(WARNING)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
}
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s]:[%s][%s]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
