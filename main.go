package main

import (
	"context"
	"database/sql"
	"example/hello/GO-KIT/account"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gokit-example/account"

	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// Create a user module that allows us to insert a user into a DB from a Web API and be able to fetch that user back using a unique user ID that we generate inside of the business logic

const dbsource = "postgresql://postgres:postgres@localhost:5432/gokitexample?sslmode=disable"

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	// Setup logger so it has timestamp, caller etc.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error
		db, err = sql.Open("postgres", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	// Parse the CLI flag is we get one
	flag.Parse()

	// Creates a non-nil empty context
	ctx := context.Background()
	var srv account.Service
	{
		// To create the microservice we first need to create a repository
		repository := account.NewRepo(db, logger)

		// Then create the service itself by passing in the repo and the logger into the call to NewService
		srv = account.NewService(respository, logger)
	}

	// Create errors channel
	errs := make(chan error)

	// Go routine to check if there's a SIG termination in the operating system
	// If that happens then will generate errors
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// Create endpoints for our service
	endpoints := account.Endpoints(srv)

	// Go routine which will execute our server
	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := account.NewHTTPServer(ctx, endpoints)
		// If given an error will be piped into our errs channel
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
