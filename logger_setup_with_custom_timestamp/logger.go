package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)
//For this purpose we're assuming you have the directory /home/yourname/logs/ and the file /home/yourname/logs/logfile.txt
//Don't forget to run
// go get "github.com/sirupsen/logrus"

var (
	homeDir= "/home/yourname"
	logDir    = filepath.Join(homeDir, "logs")
	logFile   = filepath.Join(logDir, "logfile.txt")

	mainLogger *logrus.Logger
)

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// You might want to change the timezone is you don't live in this area.
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return nil, fmt.Errorf("error loading location: %v", err)
	}
	entry.Time = entry.Time.In(loc)
	// You may wish to change (PST) to your own timezone. I had to play around with this, I kept getting strange output like 7:78PM, but this seems to work.
	timestamp := entry.Time.Format("2006-01-02 - 3:04PM (PST)")

	level := entry.Level.String()
	switch entry.Level {
	case logrus.InfoLevel:
		level = "[INFO]"
	case logrus.WarnLevel:
		level = "[WARN]"
	case logrus.ErrorLevel:
		level = "[ERROR]"
	case logrus.FatalLevel:
		level = "[FATAL]"
	case logrus.PanicLevel:
		level = "[PANIC]"
	case logrus.DebugLevel:
		level = "[DEBUG]"
	default:
		level = "[UNKNOWN]"
	}

	msg := fmt.Sprintf("%s %s %s\n", timestamp, level, entry.Message)
	return []byte(msg), nil
}

func setupLogging(logFile string) *logrus.Logger {
	mainLogger := logrus.New()

	mainLogFile, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening main log file: %v\n", err)
		os.Exit(1)
	}

	mainLogger.SetOutput(mainLogFile)
	mainLogger.SetLevel(logrus.DebugLevel)
	mainLogger.SetFormatter(&CustomFormatter{})

	return mainLogger
}


//usage
func main() {
	mainLogger = setupLogging(logFile)
	//how to use
	mainLogger.Info("Logger setup complete")
	mainLogger.Warn("This is a WARN message")
	mainLogger.Debug("This is a DEBUG message")
	mainLogger.Error("This is an ERROR message")
	mainLogger.Fatal("FATALITY message") 
	mainLogger.Panic("DO NOT REMAIN CALM; PANIC. This is a PANIC message")
	
}
