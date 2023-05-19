package divezapsugaredlogger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func performanceNotsentive() {
	logger, _ := zap.NewProduction()
	//logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugarLogger := logger.Sugar()

	sugarLogger.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "www.google.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugarLogger.Infof("Failed to fetch URL: %s", "www.google.com")
	sugarLogger.Debugln("this is debug info")
}

func performanceSentive() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "www.google.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func noOp() {
	logger := zap.NewNop()
	defer logger.Sync()
	logger.Info("My name", zap.String("name", "Zhao juan")) // no output
}

func buildLoggerFromCfg() *zap.SugaredLogger {
	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:      true,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := cfg.Build()
	if err != nil {
		logger = zap.NewNop()
	}
	return logger.Sugar()
}

// it has some problems: some extral info cannot be printed out
func runCustomLogger(isDevMode bool) {
	logger, _ := zap.NewProduction()
	if isDevMode {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()
	sugarLogger := logger.Sugar()

	sugarLogger.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "www.google.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugarLogger.Infof("Failed to fetch URL: %s", "www.google.com")
	sugarLogger.Debugln("this is debug info")
}

func Main() {
	//performanceNotsentive()
	//performanceSentive()
	//noOp()
	//runCustomLogger(true)
	runCustomLogger(false)
}
