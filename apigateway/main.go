package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/plugin/ochttp"
	"google.golang.org/grpc"
)

const (
	port = "8080"
)

type frontendServer struct {
	getNewIdSvcAddr string
	getNewIdSvcConn *grpc.ClientConn
}

func main() {

	ctx := context.Background()
	svc := new(frontendServer)

	mustMapEnv(&svc.getNewIdSvcAddr, "GET_NEW_ID_SERVICE_ADDR")

	mustConnGRPC(ctx, &svc.getNewIdSvcConn, svc.getNewIdSvcAddr)

	r := mux.NewRouter()
	r.HandleFunc("/getNewId", svc.getNewIdHandler).Methods(http.MethodPost)

	httpHandler := &ochttp.Handler{
		Propagation: &propagation.HTTPFormat{},
		Handler:     r,
	}

	log.Info().Msg("Starting manager service")

	if err := http.ListenAndServe(":8080", httpHandler); err != nil {
		log.Fatal().Err(err).Msg("Can’t start service")
	}

}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure())
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}
