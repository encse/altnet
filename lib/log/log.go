package log

import (
	"github.com/encse/altnet/lib/config"
	logrus "github.com/sirupsen/logrus"
)

var logger *logrus.Entry

func init() {
	conf := config.Get()
	log := logrus.New()
	logger = log.WithFields(logrus.Fields{"connectedFrom": conf.Connection.From})
}

func Trace(args ...interface{}) {
	logger.Log(logrus.TraceLevel, args...)
}

func Debug(args ...interface{}) {
	logger.Log(logrus.DebugLevel, args...)
}

func Print(args ...interface{}) {
	logger.Info(args...)
}

func Info(args ...interface{}) {
	logger.Log(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	logger.Log(logrus.WarnLevel, args...)
}

func Warning(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Log(logrus.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	logger.Log(logrus.FatalLevel, args...)
	logger.Logger.Exit(1)
}

func Panic(args ...interface{}) {
	logger.Log(logrus.PanicLevel, args...)
}

// Entry Printf family functions

func Tracef(format string, args ...interface{}) {
	logger.Logf(logrus.TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Logf(logrus.DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Logf(logrus.InfoLevel, format, args...)
}

func Printf(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Logf(logrus.WarnLevel, format, args...)
}

func Warningf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Logf(logrus.ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Logf(logrus.FatalLevel, format, args...)
	logger.Logger.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	logger.Logf(logrus.PanicLevel, format, args...)
}

// Entry Println family functions

func Traceln(args ...interface{}) {
	logger.Logln(logrus.TraceLevel, args...)
}

func Debugln(args ...interface{}) {
	logger.Logln(logrus.DebugLevel, args...)
}

func Infoln(args ...interface{}) {
	logger.Logln(logrus.InfoLevel, args...)
}

func Println(args ...interface{}) {
	logger.Infoln(args...)
}

func Warnln(args ...interface{}) {
	logger.Logln(logrus.WarnLevel, args...)
}

func Warningln(args ...interface{}) {
	logger.Warnln(args...)
}

func Errorln(args ...interface{}) {
	logger.Logln(logrus.ErrorLevel, args...)
}

func Fatalln(args ...interface{}) {
	logger.Logln(logrus.FatalLevel, args...)
	logger.Logger.Exit(1)
}

func Panicln(args ...interface{}) {
	logger.Logln(logrus.PanicLevel, args...)
}
