package main

// import (
// 	"flag"
// 	"fmt"
// )

// INFO: Package Main
// SOURCE: https://thenewstack.io/understanding-golang-packages/
// When you build reusable pieces of code, you will develop a package as a shared library. But when you develop executable programs, you will use the package “main” for making the package as an executable program. The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library. The main function in the package “main” will be the entry point of our executable program. When you build shared libraries, you will not have any main package and main function in the package.

import (
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"

	"flag"
	"fmt"

	// cli "github.com/behance/go-cli"

	"github.com/bossjones/go-chatbot-lab/config"
	logUtils "github.com/bossjones/go-chatbot-lab/shared/log-utils"
	"github.com/bossjones/go-chatbot-lab/shared/version"
)

// type ConfigValue string

var (
	argConfigFile   string
	argFixtureFile  string
	flagShowVersion bool
	// dbReady         chan db.Wrapper = make(chan db.Wrapper, 1)
	gLogger logUtils.Logger
)

func parseFlags() {
	flag.StringVar(&argConfigFile, "config-file", "", "configuration file")
	flag.BoolVar(&flagShowVersion, "version", false, "show version")
	flag.StringVar(&argFixtureFile, "fixtures", "", "JSON file with fixtures (see shared/helpers/fixtures for format)")
	flag.Parse()
}

// func isFile(path string) bool {
// 	_, err := os.Stat(path)
// 	return !os.IsNotExist(err)
// }

func panicGuard() {
	panicMsg := recover()
	if panicMsg != nil {
		gLogger.Error(fmt.Sprintf("PANIC CAUGHT: %+v", panicMsg))
	}

	stack := string(debug.Stack())
	if !strings.Contains(os.Getenv("CHATBOT_CLUSTER_INFO"), "dev/local") {
		stack = strings.Replace(stack, "\n", " -- ", -1)
	}
	gLogger.Error(fmt.Sprintf("BEGIN PANIC FULLSTACK: %s", stack))
	gLogger.Error("Exiting...")

	os.Exit(1)
}

func setupLogger() {
	gLogger = logUtils.NewLogger()
}

func createConfig() config.Config {
	if argConfigFile == "" {
		gLogger.Info("No config file specified. Attempting to generate from ENV")
	}
	cfg, err := config.NewConfig(argConfigFile)
	if err != nil {
		gLogger.Crit("Error reading config.", "error", err)
	}
	// For local dev environments, we print the config as pretty JSON.
	// Otherwise, we show it all on one line to be more Splunk-friendly.
	if strings.Contains(os.Getenv("CHATBOT_CLUSTER_INFO"), "dev/local") {
		gLogger.Info(fmt.Sprintf("Chatbot Config: %+v", cfg))
	} else {
		gLogger.Info("**** Chatbot Config ****", "config", cfg)
	}
	if cfg.Debug {
		logUtils.SetLevel(logUtils.DebugLevel)
	}

	return cfg
}

// func setupLogging(logLevel string, isDebug bool, logLocation string, logAppName string) {
// 	os.Setenv("LOG_APP_NAME", logAppName)
// 	log.AlwaysShowColors(true)
// 	if logLevel == "debug" || isDebug {
// 		log.SetLevel(log.DebugLevel)
// 	} else {
// 		log.SetLevel(log.InfoLevel)
// 	}

// 	log.SetFormatter(formatters.SumologicFormatter{})

// 	// set the log location, defaulting to stdout
// 	if logLocation != "stdout" {
// 		logFile, err := os.OpenFile(logLocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 		if err != nil {
// 			log.Errorf("error opening log file: %v", err)
// 		}
// 		defer logFile.Close()
// 		log.SetOutput(logFile)
// 	}
// }

// Golang application auto build versioning
// SOURCE: https://stackoverflow.com/questions/11354518/golang-application-auto-build-versioning

func main() {

	// set up signal handler
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, syscall.SIGTERM)

	defer panicGuard()

	parseFlags()

	if flagShowVersion {
		fmt.Println(version.Version)
		os.Exit(0)
	}
	setupLogger()

	// var db db.Wrapper
	cfg := createConfig()

	gLogger.Info("Config created", cfg)

	// TODO: Consider enabling this
	// logCtx := log15.Ctx{
	// 	"db_timeout_num_seconds": dbTimeoutNumSeconds,
	// 	"db_engine":              cfg.DBEngine,
	// 	"db_path":                cfg.DBPath,
	// }
	// gLogger.Info("Initializing database", logCtx)
	// go initDB(cfg)
	// select {
	// case db = <-dbReady:
	// 	gLogger.Info("Connected to database", logCtx)
	// case <-time.After(dbTimeoutNumSeconds * time.Second):
	// 	gLogger.Crit("Timeout waiting for database", logCtx)
	// 	os.Exit(1)
	// }

	// kv := kvwrapper.NewKVWrapperWithAuth(
	// 	[]string{cfg.KVServer},
	// 	kvwrapper_etcd.EtcdWrapper{},
	// 	cfg.KVUsername,
	// 	cfg.KVPassword,
	// )

	// server, err := startServer(db, kv, cfg)
	// if err != nil {
	// 	gLogger.Error(fmt.Sprintf("Cannot start server because \"%s\"", err.Error()))
	// 	os.Exit(2)
	// }

	<-sigchan
	// server.Stop()
	// for i := 0; i < numWorkers; i++ {
	// 	// signal workers
	// 	workerObjs[i].Stop()
	// 	gLogger.Info(
	// 		fmt.Sprintf("Worker %v stopped.", i),
	// 		"worker_id", workerObjs[i].WorkerID,
	// 	)
	// }

	// see ya
	gLogger.Info("Goodbye")
	os.Exit(0)

	// TODO: Delete all of this
	// ****************** OLD MAIN *******************************
	// versionFlag := flag.Bool("version", false, "Version")
	// // flag provided but not defined: -ginkgo.seed
	// // SOURCE: https://github.com/onsi/ginkgo/issues/296#issuecomment-249924522
	// flag.Parse()

	// // If defined, show value
	// if *versionFlag {
	// 	fmt.Println("Git Commit:", version.GitCommit)
	// 	fmt.Println("Version:", version.Version)
	// 	if version.VersionPrerelease != "" {
	// 		fmt.Println("Version PreRelease:", version.VersionPrerelease)
	// 	}
	// 	// Return multiple named results simply by return
	// 	// versionFlag, flag, and fmt will be returned(?)
	// 	return
	// }

	// fmt.Println("Hello.")
	// ****************** OLD MAIN *******************************
}


// FIXME: Move this into the regular main block so we can start chatting with the bot! (2/10/2018)
// SOURCE: https://github.com/coolspeed/century/blob/master/century.go
// func main() {
// 	fmt.Println("Server started.")

// 	port := "6666"
// 	listener, err_listen := net.Listen("tcp", ":" + port)
// 	if err_listen != nil {
// 		fmt.Println("Server listening failed. Exit.")
// 		os.Exit(1)
// 	}
// 	fmt.Println("Server started to listen on port " + port)

// 	chatRoom := NewChatRoom()
// 	// listen
// 	chatRoom.Listen()
// 	fmt.Println("chatRoom started to listen.")

// 	for {
// 		conn, err_ac := listener.Accept()
// 		if err_ac != nil {
// 			fmt.Println("Connection accepting failed.")
// 			conn.Close()
// 			time.Sleep(100 * time.Millisecond)
// 			continue
// 		}
// 		fmt.Println("A new connection accepted.")
// 		chatRoom.entrance <- conn  // ChatRoom.entrance: channel of connection
// 	}
// }
