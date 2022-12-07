package apimanager

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/lawmatsuyama/transactions/domain"
	log "github.com/sirupsen/logrus"
)

var srv *http.Server
var serviceName string

// StartAPI starts api server
func StartAPI(ctx context.Context, handler http.Handler, port, service string) {
	serviceName = service
	srv = &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           handler,
		ReadHeaderTimeout: time.Second * 20,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}
	domain.AddTaskCount()
	go func() {
		defer domain.DoneTask()
		log.WithField("service", serviceName).Debugf("listening to port [%s]", port)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.WithField("service", serviceName).WithError(err).Fatal("couldnt listen and serve")
		}
	}()
	domain.AddTaskCount()
	go func() {
		defer domain.DoneTask()
		<-ctx.Done()
		log.WithField("service", serviceName).Debug("context was canceled, shutting down server")
		gracefulStop()
	}()
}

// ShutdownAPI closes all connections and shutdown api server
func ShutdownAPI() {
	err := srv.Close()
	if err != nil {
		log.WithField("service", serviceName).WithError(err).Error("couldnt close server")
	}
}

func gracefulStop() {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	log.WithField("service", serviceName).Info("gracefully shutdowns API server")
	err := srv.Shutdown(ctxTimeout)
	if err != nil {
		log.WithField("service", serviceName).WithError(err).Errorf("couldnt gracefulStop server API")
	}
	log.WithField("service", serviceName).Debug("the server is down, good bye")
}
