# golog
Package golog is a very basic logging library that wraps the standard Go fmt package. 
golog provides basic leveled and structured logging.

# Installation
```
go get -u https://github.com/JacyGao/golog.git
```

# Example
Initialise a logger instance
```
logger := golog.New(golog.INFO)
```
For simple logging
```
logger.Info("some log messages...")
```
For formatted logging
```
logger.Infof("some %s messages...", "log")
```
For structured logging
```
logger.Infow("some log messages with tags...", "key", "value", "key2", "value2")
```
To spawn a child logger
```
childLogger := logger.With("logger", "childLogger")
childLogger.Info("some log messages...") // will log [{logger childLogger}]
```
