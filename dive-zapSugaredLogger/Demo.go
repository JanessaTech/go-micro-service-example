package divezapsugaredlogger

import (
	"time"

	"go.uber.org/zap"
)

// we use logger.Sugar() if the performance is not critical
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

// we use logger directly if the performance is critical
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

// don't output anything. This is for something going wrong when we create zap.Logger
func noOp() {
	logger := zap.NewNop()
	defer logger.Sync()
	logger.Info("My name", zap.String("name", "Zhao juan")) // no output
}

// we use a variable isDevMode to decide which type of logger we will create: for development or for production
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

// add lable to sugarLogger
func runCustomLoggerWith(isDevMode bool) {
	logger, _ := zap.NewProduction()
	if isDevMode {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()
	sugarLogger := logger.Sugar()
	sugarLogger = sugarLogger.With("requestId", "0000000000000000000")

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
	//runCustomLogger(false)
	runCustomLoggerWith(true)
}
