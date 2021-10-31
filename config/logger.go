// Package config
//created by lilei at 2021/10/7
package config

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"sync"
)

var LoggerWriterList []*loggerWriter

type loggerContent struct {
	content string
	level   log.Level
}

type loggerWriter struct {
	logFileName     string
	isWriteToStdout bool
	logger          *log.Logger
	waitGroup       *sync.WaitGroup
	consumerChannel *chan loggerContent
}

func NewLoggerWriter(level log.Level, logFileName string, isWriteToStout bool) *loggerWriter {
	// 日志目录
	logDirPath := path.Join("logs", logFileName)
	err := os.MkdirAll(logDirPath, 0644)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
	fileName := path.Join("logs", logFileName)

	// 写入文件
	src, err := rotatelogs.New(
		path.Join(logDirPath, fileName+"-%Y%m%d"),
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(3650),
	)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	// 实例化
	logger := log.New()
	logger.SetLevel(level)

	// 设置输出
	if isWriteToStout {
		writers := []io.Writer{src, os.Stdout}
		fileAndStdoutWriter := io.MultiWriter(writers...)
		logger.SetOutput(fileAndStdoutWriter)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	cc := make(chan loggerContent, 100)
	w := loggerWriter{
		logFileName, isWriteToStout, logger, &wg, &cc,
	}
	go func() {
		for lc := range cc {
			logger.Logln(lc.level, lc.content)
		}
		wg.Done()
	}()
	LoggerWriterList = append(LoggerWriterList, &w)
	return &w
}

func (w loggerWriter) Stop() {
	close(*w.consumerChannel)
	w.waitGroup.Wait()
}

func (w loggerWriter) Log(level log.Level, content string) {
	*w.consumerChannel <- loggerContent{content, level}
}

func (w loggerWriter) Info(content string) {
	w.Log(log.InfoLevel, content)
}

func (w loggerWriter) Warn(content string) {
	w.Log(log.WarnLevel, content)
}

func (w loggerWriter) Error(content string) {
	w.Log(log.ErrorLevel, content)
}

func (w loggerWriter) Fatal(content string) {
	w.Log(log.FatalLevel, content)
}

func (w loggerWriter) Panic(content string) {
	w.Log(log.PanicLevel, content)
}
