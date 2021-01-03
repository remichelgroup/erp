package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "article",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "sercice ended")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "Ressource_Management",
		Subsystem: "ArticleService",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "Ressource_Management",
		Subsystem: "ArticleService",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "Ressource_Management",
		Subsystem: "ArticleService",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	flag.Parse()

	db, err := gorm.Open(sqlite.Open("article.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var svc Service
	{
		repository := NewRepository(logger, db)
		svc = NewService(repository, logger)
	}
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}
	svc.MigrateRepo()
	endpoints := MakeEndpoints(svc)
	handler := NewHTTPServer(ctx, endpoints)
	gohandlers.CORS()
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000", "http://administration.remichel-cc.com"}))
	s := &http.Server{
		Addr:         ":9091",
		Handler:      ch(handler),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	errs := make(chan error)
	go func() {
		errs <- s.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)

	sig := <-signalChan
	level.Error(logger).Log("Recieved terminate, graceful shutdown", sig)
	s.Shutdown(ctx)
}
