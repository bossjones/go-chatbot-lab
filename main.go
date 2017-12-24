package main

// import (
// 	"flag"
// 	"fmt"
// )

import (
	"os"

	"flag"
	"fmt"

	// cli "github.com/behance/go-cli"
	"github.com/behance/go-logging/formatters"
	"github.com/behance/go-logging/log"
)

type ConfigValue string

func setupLogging(logLevel string, isDebug bool, logLocation string, logAppName string) {
	os.Setenv("LOG_APP_NAME", logAppName)
	log.AlwaysShowColors(true)
	if logLevel == "debug" || isDebug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.SetFormatter(formatters.SumologicFormatter{})

	// set the log location, defaulting to stdout
	if logLocation != "stdout" {
		logFile, err := os.OpenFile(logLocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Errorf("error opening log file: %v", err)
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	}
}

func main() {

	versionFlag := flag.Bool("version", false, "Version")
	// flag provided but not defined: -ginkgo.seed
	// SOURCE: https://github.com/onsi/ginkgo/issues/296#issuecomment-249924522
	flag.Parse()

	if *versionFlag {
		fmt.Println("Git Commit:", GitCommit)
		fmt.Println("Version:", Version)
		if VersionPrerelease != "" {
			fmt.Println("Version PreRelease:", VersionPrerelease)
		}
		// Return multiple named results simply by return
		// versionFlag, flag, and fmt will be returned(?)
		return
	}

	fmt.Println("Hello.")
}
