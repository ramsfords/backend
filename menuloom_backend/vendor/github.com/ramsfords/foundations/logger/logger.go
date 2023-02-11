package logger

import (
	"log"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

// New constructs a Sugared Logger that writes to stdout and
// provides human-readable timestamps.
func New(serviceName ...string) Logger {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"} //[]string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	if len(serviceName) != 0 {
		config.InitialFields = map[string]interface{}{
			"service": serviceName,
		}
	} else {
		config.InitialFields = map[string]interface{}{
			"service": serviceName,
		}
	}
	logs, err := config.Build(zap.AddCaller(), zap.Fields())
	if err != nil {
		err = errors.Wrap(err, "error initiating loggers")
		log.Fatal(err)
	}
	return Logger{logs.Sugar()}

}
