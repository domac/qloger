package qloger

import (
	"github.com/Sirupsen/logrus"
	"os"
)

//统一日志接口
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})
}

//带循环日志功能logger
func NewRotatorQLogger(outString, loglevel string,
	openRotate, rotateDaily, IsCompress bool,
	maxSize int64) (Logger, error) {

	defaultLogger := logrus.New()
	err := SetLogOut(defaultLogger, outString, openRotate, rotateDaily, IsCompress, maxSize)
	if err != nil {
		return nil, err
	}
	SetLogLevel(defaultLogger, loglevel)
	defaultLogger.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	}
	return defaultLogger, nil
}

//默认logger
func NewQLogger(outString, loglevel string) (Logger, error) {
	return NewRotatorQLogger(outString, loglevel, false, false, false, -1)
}

//设置日志输出方式
func SetLogOut(log *logrus.Logger, outString string,
	openRotate, rotateDaily, IsCompress bool, maxSize int64) error {
	switch outString {
	case "stdout":
		log.Out = os.Stdout
	case "stderr":
		log.Out = os.Stderr
	default:
		//是否开启循环日志
		if openRotate {
			opts := &Options{
				RotateDaily: rotateDaily,
				Compress:    IsCompress,
				MaximumSize: maxSize,
			}
			writer, err := NewWriter(outString, opts)
			if err != nil {
				return err
			}
			log.Out = writer
		} else {
			f, err := os.OpenFile(outString, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				return err
			}
			log.Out = f
		}
	}
	return nil
}

//设置日志级别
func SetLogLevel(log *logrus.Logger, levelString string) error {
	level, err := logrus.ParseLevel(levelString)
	if err != nil {
		return err
	}
	log.Level = level
	return nil
}
