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
var Conf *Config

type Config struct {
	Prefix     string
	LogDir     string
	Level      int
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

func GetInfoLogger() *log.Logger {
	return infoLog
}

// InitConfig ...
// 5=debug,10=warn,15=error
func InitConfig(_config *Config) io.Writer {
	Conf = _config
	if len(Conf.LogDir) <= 0 {
		Conf.LogDir = "./logs/"
	}
	if Conf.Level <= 0 {
		Conf.Level = 5
	}
	if Conf.MaxBackups <= 0 {
		Conf.MaxBackups = 3
	}
	if Conf.MaxSize <= 0 {
		Conf.MaxSize = 10
	}
	if Conf.MaxAge <= 0 {
		Conf.MaxAge = 7
	}
	_level = Conf.Level
	errorLogFile := &Logger{
		Filename:   Conf.LogDir + "error.log",
		MaxSize:    Conf.MaxSize, // megabytes
		MaxBackups: Conf.MaxBackups,
		MaxAge:     Conf.MaxAge, // days
	}
	errorLog = log.New(io.MultiWriter(os.Stdout, errorLogFile), "["+Conf.Prefix+"] [ERROR] ", log.Ldate|log.Lmicroseconds|log.Llongfile)

	infoLogFile := &Logger{
		Filename:   Conf.LogDir + "info.log",
		MaxSize:    Conf.MaxSize, // megabytes
		MaxBackups: Conf.MaxBackups,
		MaxAge:     Conf.MaxAge, // days
	}
	writer := io.MultiWriter(os.Stdout, infoLogFile)
	infoLog = log.New(writer, "["+Conf.Prefix+"] [INFO] ", log.Ldate|log.Lmicroseconds|log.Llongfile)

	warnLogFile := &Logger{
		Filename:   Conf.LogDir + "warn.log",
		MaxSize:    Conf.MaxSize, // megabytes
		MaxBackups: Conf.MaxBackups,
		MaxAge:     Conf.MaxAge, // days
	}
	warnLog = log.New(io.MultiWriter(os.Stdout, warnLogFile), "["+Conf.Prefix+"] [WARN] ", log.Ldate|log.Lmicroseconds|log.Llongfile)

	return writer
}
func ErrorSkip(format string, v ...interface{}) {
	if _level == -999 {
		fmt.Printf("Log Uninitialized:"+format, v...)
	} else {
		if _level <= 15 {
			errorLog.Output(2, fmt.Sprintf(format, v...))
		}
	}
}
func Error(format string, v ...interface{}) {
	if _level == -999 {
		fmt.Printf("Log Uninitialized:"+format, v...)
	} else {
		if _level <= 15 {
			errorLog.Output(2, fmt.Sprintf(format, v...))
			errorLog.Output(2, fmt.Sprintf("%s", Stack(5)))
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

func Stack(skip int) []byte {
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
	data = append(data[:1], data[skip:]...)
	return bytes.Join(data, line)
}
