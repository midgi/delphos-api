package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/migdi/delphos-api/handlers"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/ifrit/sigmon"
)

var listenAddress = flag.String(
	"listenAddress",
	"0.0.0.0:8080",
	"The host:port that the server is bound to.",
)

var logLevel = flag.String(
	"logLevel",
	"info",
	"log level: debug, info, error or fatal",
)

func main() {
	logger := configureLogger()
	flag.Parse()

	validateFlags(logger)

	router := configureRouter(logger)

	var server ifrit.Runner
	server = http_server.New(*listenAddress, router)
	members := grouper.Members{
		{"server", server},
	}

	group := grouper.NewOrdered(os.Interrupt, members)

	monitor := ifrit.Invoke(sigmon.New(group))
	logger.Info("started")

	err := <-monitor.Wait()
	if err != nil {
		logger.Error("exited-with-failure", err)
		os.Exit(1)
	}

	logger.Info("exited")
}

func configureRouter(logger lager.Logger) http.Handler {
	infoHandler := handlers.NewInfoHandler(logger)

	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/user", UserHandler)
	router.Handle("/info", infoHandler)
	return router
}

func configureLogger() lager.Logger {
	var lagerLogLevel lager.LogLevel
	switch *logLevel {
	case "debug":
		lagerLogLevel = lager.DEBUG
	case "info":
		lagerLogLevel = lager.INFO
	case "error":
		lagerLogLevel = lager.ERROR
	case "fatal":
		lagerLogLevel = lager.FATAL
	default:
		panic(fmt.Errorf("unknown log level: %s", logLevel))
	}

	logger := lager.NewLogger("delphos")
	sink := lager.NewReconfigurableSink(lager.NewWriterSink(os.Stdout, lager.DEBUG), lagerLogLevel)
	logger.RegisterSink(sink)

	return logger
}

func validateFlags(logger lager.Logger) {
	_, portString, err := net.SplitHostPort(*listenAddress)
	if err != nil {
		logger.Fatal("failed-invalid-listen-address", err)
	}
	_, err = net.LookupPort("tcp", portString)
	if err != nil {
		logger.Fatal("failed-invalid-listen-port", err)
	}
}
