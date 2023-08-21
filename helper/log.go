package helper

import (
	"io"
	"log"
	"os"
	"sync"
	"time"
)

/**
自己实现一个日志库。
	1、 写文件
	2、 每到1000条  进行数据刷盘、  每秒进行数据刷盘    defer 刷盘
	3、 日志切割：  100M进行切割、
	4、 文件路径、  年/月/天-xxx.log
	5、 提供Write方法  满足需要立即写入日志的需求
	6、 提供日志级别  info  error   error 需要打印堆栈、 并能发送到钉钉
*/

var Logger *MyLogger

// MyLogger logger封装
type MyLogger struct {
	infoLogger
	errorLogger
}

func InitLogger() {

	infoFile := createInfoFile()
	errorFile := createErrorFile()

	Logger = &MyLogger{
		infoLogger{
			queue:        make(chan string, 1050),
			logger:       log.New(infoFile, "", log.LstdFlags),
			currencyFile: infoFile,
			conf:         initLogConfig(),
		},
		errorLogger{
			queue:        make(chan string, 1050),
			logger:       log.New(io.MultiWriter(errorFile, os.Stderr), "", log.LstdFlags),
			currencyFile: errorFile,
			conf:         initLogConfig(),
		},
	}
	go Logger.timerFlush()
}

func (l *MyLogger) Info(message string) {

	l.infoLogger.queue <- message
	if len(l.infoLogger.queue) > 1000 {
		go l.infoLogger.flush()
	}
}
func (l *MyLogger) Write(message string) {
	l.infoLogger.queue <- message
	l.infoLogger.flush()
}

func (l *MyLogger) Error(message string) {
	l.errorLogger.queue <- message
	if len(l.errorLogger.queue) > 1000 {
		go l.errorLogger.flush()
	}
}

func (l *MyLogger) timerFlush() {
	timer := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-timer.C:
			l.Info("日志刷盘")
			l.infoLogger.flush()
			l.errorLogger.flush()
		default:
			time.Sleep(time.Second * 1)
		}
	}
}

//日志配置
type logConfig struct {
	maxSize int //单位 M
}

func initLogConfig() logConfig {
	subViper := ViperConfig.Sub("log")
	logSize := subViper.GetInt("size")
	if logSize == 0 {
		logSize = 100
	}

	return logConfig{
		maxSize: logSize,
	}
}

func createInfoFile() *os.File {
	currentDay := time.Now().Format("20060102")
	filePathDir := "./runtime/log/" + currentDay
	err := os.MkdirAll(filePathDir, os.ModePerm)
	if err != nil {
		log.Fatal("日志目录创建失败、 请检查权限", err.Error())
	}

	file, err := os.OpenFile(filePathDir+"/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("日志文件创建失败、 请检查权限", err.Error())
	}
	return file
}

func createErrorFile() *os.File {
	currentDay := time.Now().Format("20060102")
	filePathDir := "./runtime/log/" + currentDay
	err := os.MkdirAll(filePathDir, os.ModePerm)
	if err != nil {
		log.Fatal("日志目录创建失败、 请检查权限", err.Error())
	}

	file, err := os.OpenFile(filePathDir+"/error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("日志文件创建失败、请检查权限", err.Error())
	}
	return file
}

//info日志 struct 提供数据buffer  channel
type infoLogger struct {
	currencyFile *os.File
	locker       sync.Mutex
	queue        chan string
	logger       *log.Logger
	conf         logConfig
}

func (il *infoLogger) flush() {
	il.locker.Lock()
	defer il.locker.Unlock()

	//双重检查
	queueLen := len(il.queue)
	if queueLen == 0 {
		return
	}

	//数据刷盘
	for i := 0; i < queueLen; i++ {
		il.logger.Println(<-il.queue)
	}

	//日志切割
	if fileSize, _ := il.currencyFile.Stat(); fileSize.Size() > int64(il.conf.maxSize*1024*1024) {
		_ = il.currencyFile.Close()
		err := os.Rename(il.currencyFile.Name(), il.currencyFile.Name()+"."+time.Now().Format("150405"))
		if err != nil {
			log.Println(err.Error())
		}

		//重新打开一个文件句柄、  应该比文件复制 然后清空原文件要快
		newInfoFile := createInfoFile()
		il.logger = log.New(newInfoFile, "", log.LstdFlags)
		il.currencyFile = newInfoFile
	}

}

//error 日志 提供数据buffer channel
type errorLogger struct {
	currencyFile *os.File
	locker       sync.Mutex
	queue        chan string
	logger       *log.Logger
	conf         logConfig
}

func (el *errorLogger) flush() {
	el.locker.Lock()
	defer el.locker.Unlock()

	//双重检查
	if len(el.queue) == 0 {
		return
	}

	//数据刷盘  日志切割
	for msg := range el.queue {
		el.logger.Println(msg)
	}

	if fileSize, _ := el.currencyFile.Stat(); fileSize.Size() > int64(el.conf.maxSize*1024*1024) {
		_ = el.currencyFile.Close()
		_ = os.Rename(el.currencyFile.Name(), el.currencyFile.Name()+"."+time.Now().Format("150405"))

		//重新打开一个文件句柄、  应该比文件复制 然后清空原文件要快
		newErrorFile := createErrorFile()
		el.logger = log.New(newErrorFile, "", log.LstdFlags)
		el.currencyFile = newErrorFile
	}

}
