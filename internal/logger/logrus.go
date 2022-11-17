package logger

import (
	"bufio"
	"fmt"
	"github.com/Songguangyun/go-web/internal/global"
	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func InitLogger(moduleName string) *logrus.Logger {
	var logger = logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(src)
	logger.SetOutput(writer)
	hook := NewLfsHook(time.Duration(global.ConfigServer.Log.RotationTime)*time.Hour, uint(global.ConfigServer.Log.RemainRotationCount), moduleName)
	logger.AddHook(hook)
	logger.SetFormatter(getLogFormatter(true))
	logger.SetReportCaller(true)
	return logger
}

// 获取日志格式
func getLogFormatter(isConsole bool) *nested.Formatter {
	formatter := &nested.Formatter{
		HideKeys:        false,
		TimestampFormat: "2006-01-02 15:04:05.000",
		FieldsOrder:     []string{"PID", "FilePath"},
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			funcInfo := runtime.FuncForPC(frame.PC)
			if funcInfo == nil {
				return "error during runtime.FuncForPC"
			}
			fullpath, line := funcInfo.FileLine(frame.PC)
			return fmt.Sprintf(" [%v:%v]", filepath.Base(fullpath), line)
		},
	}
	if isConsole {
		formatter.NoColors = false
	} else {
		formatter.NoColors = true
	}

	return formatter
}

// NewLfsHook 设置钩子
func NewLfsHook(rotationTime time.Duration, maxRemainNum uint, moduleName string) logrus.Hook {
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: initRotateLogs(rotationTime, maxRemainNum, "all", moduleName),
		logrus.InfoLevel:  initRotateLogs(rotationTime, maxRemainNum, "all", moduleName),
		logrus.WarnLevel:  initRotateLogs(rotationTime, maxRemainNum, "all", moduleName),
		logrus.ErrorLevel: initRotateLogs(rotationTime, maxRemainNum, "all", moduleName),
	}, getLogFormatter(false))
	return lfsHook
}

func initRotateLogs(rotationTime time.Duration, maxRemainNum uint, level string, moduleName string) *rotatelogs.RotateLogs {
	if moduleName != "" {
		moduleName = moduleName + "."
	}
	writer, err := rotatelogs.New(
		global.ConfigServer.Log.StorageLocation+moduleName+level+"."+"%Y-%m-%d"+".log",
		rotatelogs.WithRotationTime(rotationTime),  // 日志周期
		rotatelogs.WithRotationCount(maxRemainNum), // 只保留最近的N个日志文件
	)
	if err != nil {
		panic(err)
	} else {
		return writer
	}
}

func NewInfo(args ...interface{}) {
	global.Logger.WithFields(logrus.Fields{
		"PID": os.Getpid(),
	}).Infoln(args)
}
func NewWarn(args ...interface{}) {
	global.Logger.WithFields(logrus.Fields{
		"PID": os.Getpid(),
	}).Warnln(args)
}
func NewDebug(args ...interface{}) {
	global.Logger.WithFields(logrus.Fields{
		"PID": os.Getpid(),
	}).Debugln(args)
}

func NewError(args ...interface{}) {
	global.Logger.WithFields(logrus.Fields{
		"PID": os.Getpid(),
	}).Errorln(args)
}
