package main

import (
	"context"
	"github.com/farwydi/cleanwhale/config"
	"github.com/farwydi/cleanwhale/log"
	"github.com/farwydi/cleanwhale/metrics"
	"github.com/farwydi/cleanwhale/transport/whalehttp"
	"github.com/farwydi/cleanwhale/wave"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	var cfg AppConfig
	err := config.LoadConfigs(&cfg, "static/config.yml", "config.yml")
	if err != nil {
		log.Fatal(err)
	}

	logger, err := log.NewLogger(cfg.Project)
	if err != nil {
		log.Fatal(err)
	}

	mlog := logger.Named("main")

	log.RegisterLogger(logger)
	metrics.RegisterMetrics("", mlog)

	app, cleanup, err := setup(cfg)
	if err != nil {
		mlog.Fatal("Fail setup",
			zap.Error(err))
	}

	defer cleanup()

	sig := wave.NewWave(context.Background(), mlog)

	sig.AddServer(whalehttp.NewHTTPServer(
		cfg.GraphQLTransport, logger.Named("web"), app.http))

	sig.Run()
}

func newApplication(http http.Handler) application {
	return application{
		http: http,
	}
}

type application struct {
	http http.Handler
}
