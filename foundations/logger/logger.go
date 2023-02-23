package logger

import (
	"io"
	"log/syslog"
	"os"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/ramsfords/backend/configs"
	"github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
)

var levelsMap = map[log.Lvl]logrus.Level{
	log.DEBUG: logrus.DebugLevel,
	log.INFO:  logrus.InfoLevel,
	log.WARN:  logrus.WarnLevel,
	log.ERROR: logrus.ErrorLevel,
	log.OFF:   logrus.FatalLevel,
}

type AppLogger struct {
	*logrus.Logger
	*newrelic.Application
	prefix   string
	level    log.Lvl
	MsgField string
}

var Logger = AppLogger{}

func New(conf *configs.Config, prefix string) (*AppLogger, error) {
	logg := logrus.New()
	hook, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")
	if err != nil {
		return nil, err
	}
	logg.ReportCaller = true
	logg.AddHook(hook)
	logg.Trace("Starting logger")
	logg.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})
	Logger = AppLogger{
		Logger:   logg,
		prefix:   prefix,
		MsgField: "",
	}
	switch logg.GetLevel() {
	case logrus.DebugLevel, logrus.TraceLevel:
		Logger.level = log.DEBUG
	case logrus.InfoLevel:
		Logger.level = log.INFO
	case logrus.WarnLevel:
		Logger.level = log.WARN
	case logrus.ErrorLevel:
		Logger.level = log.ERROR
	case logrus.FatalLevel, logrus.PanicLevel:
		Logger.level = log.OFF
	default:
		Logger.level = log.INFO
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("firstshipper-api"),
		newrelic.ConfigEnabled(true),
		newrelic.ConfigLicense(conf.NewRelic.License),
		newrelic.ConfigDebugLogger(os.Stdout),
		func(config *newrelic.Config) {
			logrus.SetLevel(logrus.DebugLevel)
			config.Logger = Logger
		},
	)

	if err != nil {
		return nil, err
	}
	Logger.Application = app
	return &Logger, nil

}

func (appLogger AppLogger) Error(msg string, context map[string]interface{}) {
	appLogger.Errorf(msg, context)
	appLogger.Application.RecordLog(newrelic.LogData{
		Message:   msg,
		Timestamp: time.Now().Unix(),
		Severity:  "error",
	})
	// strEvent := fmt.Sprintf("my own event: %v", context)
	// fmt.Println(strEvent)
	// appLogger.Application.RecordCustomEvent("my own event", map[string]interface{}{"error": msg, "context": strEvent})
}

func (appLogger AppLogger) Warn(msg string, context map[string]interface{}) {
	appLogger.WithFields(context).Warn(msg)
	appLogger.Application.RecordLog(newrelic.LogData{
		Message:   msg,
		Timestamp: time.Now().Unix(),
		Severity:  "warn",
	})
}

func (appLogger AppLogger) Info(msg string, context map[string]interface{}) {
	appLogger.WithFields(context).Info(msg)
	appLogger.Application.RecordLog(newrelic.LogData{
		Message:   msg,
		Timestamp: time.Now().Unix(),
		Severity:  "info",
	})
}

func (appLogger AppLogger) Debug(msg string, context map[string]interface{}) {
	appLogger.Application.RecordLog(newrelic.LogData{
		Message:   msg,
		Timestamp: time.Now().Unix(),
		Severity:  "debug",
	})
}
func (appLogger AppLogger) DebugEnabled() bool {
	return true
}
func Error(err interface{}, message string) {
	Logger.Error(message, map[string]interface{}{"error": err})
}
func Info(data interface{}, message string) {
	Logger.Info(message, map[string]interface{}{"info": data})
}

func (l *AppLogger) fields(j log.JSON) (string, logrus.Fields) {
	msg, _ := j[l.MsgField].(string)
	delete(j, l.MsgField)
	return l.Prefix() + msg, logrus.Fields(j)
}

func (l *AppLogger) SetHeader(h string) {
}

func (l *AppLogger) Output() io.Writer {
	return l.Out
}

func (l *AppLogger) SetOutput(w io.Writer) {
	l.Logger.SetOutput(w)
}

func (l *AppLogger) Prefix() string {
	return l.prefix
}

func (l *AppLogger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *AppLogger) Level() log.Lvl {
	return l.level
}

func (l *AppLogger) SetLevel(lvl log.Lvl) {
	l.level = lvl
	l.Logger.SetLevel(levelsMap[lvl])
}

func (l *AppLogger) Printj(j log.JSON) {
	msg, fields := l.fields(j)
	l.WithFields(fields).Print(msg)
}

func (l *AppLogger) Debugj(j log.JSON) {
	msg, fields := l.fields(j)
	l.WithFields(fields).Debug(msg)
}

func (l *AppLogger) Infoj(j log.JSON) {
	msg, fields := l.fields(j)
	l.WithFields(fields).Info(msg)
}
