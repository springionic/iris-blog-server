// created by lilei at 2021/10/7
package server

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"sync"
)

//type AccessLog struct {
//	StartTime   time.Time
//	EndTime     time.Time
//	StatusCode  int
//	LatencyTime time.Duration
//	ClientIP    string
//	Svc         string
//	Uid         string
//	ReqMethod   string
//	ReqUri      string
//	ReqBody     []byte
//	RspBody     []byte
//}

type loggerContent struct {
	content string
	level   log.Level
}

type AccessLogWriter struct {
	logFileName     string
	isWriteToStdout bool
	logger          *log.Logger
	waitGroup       *sync.WaitGroup
	consumerChannel *chan loggerContent
}

func NewAccessLogWriter(level log.Level, logFileName string, isWriteToStout bool) *AccessLogWriter {
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
	w := AccessLogWriter{
		logFileName, isWriteToStout, logger, &wg, &cc,
	}
	go func() {
		for lc := range cc {
			logger.Logln(lc.level, lc.content)
		}
		wg.Done()
	}()
	return &w
}

func (w AccessLogWriter) Log(level log.Level, content string) {
	*w.consumerChannel <- loggerContent{content, level}
}

func (w AccessLogWriter) Info(content string) {
	w.Log(log.InfoLevel, content)
}

func (w AccessLogWriter) Warn(content string) {
	w.Log(log.WarnLevel, content)
}

func (w AccessLogWriter) Error(content string) {
	w.Log(log.ErrorLevel, content)
}

func (w AccessLogWriter) Fatal(content string) {
	w.Log(log.FatalLevel, content)
}

func (w AccessLogWriter) Panic(content string) {
	w.Log(log.PanicLevel, content)
}
