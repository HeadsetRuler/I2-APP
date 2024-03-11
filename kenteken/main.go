package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/sirupsen/logrus"
)

type configuration struct {
	AllowedPlates []string
	logFile string
	logLevel string
}

var log = logrus.New()


func main() {
	configPath := flag.String("config", "config.json", "path to the configuration file")
	logLevel := flag.StringP("loglevel", "l", "", "log level")
	flag.Parse()
	config, configErr := loadConfig(*configPath)

	logFile, logFileErr := os.OpenFile(config.logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if logFileErr != nil {
		log.WithError(logFileErr).WithField("path", config.logFile).Warn("Failed to open log file, using stderr")
	} else {
		log.Out = logFile
		log.SetNoLock()
	}
	if configErr != nil {
		log.WithError(configErr).WithField("path", *configPath).Warn("Failed to parse config, using defaults")
	} else {
		log.WithField("path", *configPath).Info("Loaded config")
	}
	if len(*logLevel) > 0 {
		config.logLevel = *logLevel
	}
	if len(config.logLevel) == 0{
		level, logLevelErr := logrus.ParseLevel(config.logLevel)
		if logLevelErr != nil {
			log.WithError(logLevelErr).WithField("level", config.logLevel).Warn("Failed to parse log level from config, using default (error)")
		} else {
			log.SetLevel(level)
		}
	}
 

	kentekens := make(map[string]bool)
	for _, p := range config.AllowedPlates {
		kentekens[p] = true
	}

	plate := flag.Arg(0)
	
	now := time.Now()
	if len(plate) == 0 {
		fmt.Print("Kenteken: ")
		n, err := fmt.Scan(&plate)
		if n == 0 || err != nil {
			fmt.Println("Geen kenteken opgegeven")
			return
		}
	}
	if kentekenChecker(plate, kentekens) {
		fmt.Println(groet(now))
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
	}
}

// Loads and parses config, setting defaults if necessary
func loadConfig(path string) (configuration, error) {
	var config configuration
	file, err := os.Open(path)
	if err != nil {
		log.WithError(err).WithField("path", path).Warn("Failed to open config file, using defaults")
		err = nil
	} else {
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
	}
	
	// Set defaults
	if err != nil || len(config.logFile) == 0 {
		config.logFile = "kenteken.log"
	}
	if(err != nil || len(config.AllowedPlates) == 0) {
		config.AllowedPlates = []string{"12-AB-34", "56-CD-78", "90-EF-12"}
	}
	if(err != nil || len(config.logLevel) == 0) {
		config.logLevel = "error"
	}
	return config, err
}

func groet(now time.Time) string {
	const phrase = "! Welkom bij Fonteyn Vakantieparken"
	hour := now.Hour()
	switch {
	case hour < 7:
		return "Sorry, de parkeerplaats is ’s nachts gesloten"
	case hour < 12:
		return fmt.Sprintf("Goedemorgen%v", phrase)
	case hour < 18:
		return fmt.Sprintf("Goedemiddag%v", phrase)
	case hour < 23:
		return fmt.Sprintf("Goedenavond%v", phrase)
	default:
		return "Sorry, de parkeerplaats is ’s nachts gesloten"
	}
}

func kentekenChecker(plate string, kentekens map[string]bool) bool {
	allowed := kentekens[plate]
	log.WithField("plate", plate).WithField("allowed", allowed).Info("Checking license plate")
	return allowed
}
