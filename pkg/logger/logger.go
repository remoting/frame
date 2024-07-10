package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

var errorLog, infoLog, warnLog *log.Logger
var _level int = -999

type Config struct {
	Prefix     string
	LogDir     string
	Level      int
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

// InitConfig ...
// 5=debug,10=warn,15=error
func InitConfig(conf Config) io.Writer {
	if len(conf.LogDir) <= 0 {
		conf.LogDir = "./logs/"
	}
	if conf.Level <= 0 {
		conf.Level = 5
	}
	if conf.MaxBackups <= 0 {
		conf.MaxBackups = 3
	}
	if conf.MaxSize <= 0 {
		conf.MaxSize = 10
	}
	if conf.MaxAge <= 0 {
		conf.MaxAge = 7
	}
	_level = conf.Level
	errorLogFile := &Logger{
		Filename:   conf.LogDir + "error.log",
		MaxSize:    conf.MaxSize, // megabytes
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge, // days
	}
	errorLog = log.New(io.MultiWriter(os.Stdout, errorLogFile), "["+conf.Prefix+"] [ERROR] ", log.Ldate|log.Lmicroseconds|log.Llongfile)

	infoLogFile := &Logger{
		Filename:   conf.LogDir + "info.log",
		MaxSize:    conf.MaxSize, // megabytes
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge, // days
	}
	writer := io.MultiWriter(os.Stdout, infoLogFile)
	infoLog = log.New(writer, "["+conf.Prefix+"] [INFO] ", log.Ldate|log.Lmicroseconds|log.Llongfile)

	warnLogFile := &Logger{
		Filename:   conf.LogDir + "warn.log",
		MaxSize:    conf.MaxSize, // megabytes
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge, // days
	}
	warnLog = log.New(io.MultiWriter(os.Stdout, warnLogFile), "["+conf.Prefix+"] [WARN] ", log.Ldate|log.Lmicroseconds|log.Llongfile)

	return writer
}

func Error(format string, v ...interface{}) {
	if _level == -999 {
		fmt.Printf("Log Uninitialized:"+format, v...)
	} else {
		if _level <= 15 {
			errorLog.Output(2, fmt.Sprintf(format, v...))
			errorLog.Output(2, fmt.Sprintf("%s", Stack()))
		}
	}
}
func Info(format string, v ...interface{}) {
	if _level == -999 {
		fmt.Printf("Log Uninitialized:"+format, v...)
	} else {
		if _level <= 5 {
			infoLog.Output(2, fmt.Sprintf(format, v...))
		}
	}
}
func Warn(format string, v ...interface{}) {
	if _level == -999 {
		fmt.Printf("Log Uninitialized:"+format, v...)
	} else {
		if _level <= 10 {
			warnLog.Output(2, fmt.Sprintf(format, v...))
		}
	}
}

func Stack() []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			buf = buf[:n]
			break
		}
		buf = make([]byte, 2*len(buf))
	}
	line := []byte("\n")
	data := bytes.Split(buf, line)
	data = append(data[:1], data[5:]...)
	return bytes.Join(data, line)
}
