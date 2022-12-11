package apimanager_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lawmatsuyama/transactions/domain"
	"github.com/lawmatsuyama/transactions/infra/apimanager"

	"github.com/stretchr/testify/assert"
)

func TestApiManagerShouldStartAndShutdownOk(t *testing.T) {
	t.Run("01_should_start_and_shutdown_ok", func(t *testing.T) {
		if !*integration {
			return
		}
		defer t.Cleanup(func() {
			domain.CleanupWaitGroup()
		})
		r := chi.NewRouter()
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

		apimanager.StartAPI(context.TODO(), r, "8888", "test")
		time.Sleep(time.Second * 5)
		b, err := GetOk(8888)

		assert.Nil(t, err, "http get should return ok")
		assert.Equal(t, string(b), "ok", "response body should return ok")

		apimanager.ShutdownAPI()
		time.Sleep(time.Second * 5)
		b, err = GetOk(8888)
		assert.Nil(t, b)
		assert.NotNil(t, err, "http get should return error")
	})

}

func TestApiManagerShouldStartAndGracefulShutdownOk(t *testing.T) {
	t.Run("02_should_start_and_graceful_shutdown_ok", func(t *testing.T) {
		if !*integration {
			return
		}
		defer t.Cleanup(func() {
			domain.CleanupWaitGroup()
		})
		r := chi.NewRouter()
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

		ctx, cancel := context.WithCancel(context.TODO())
		apimanager.StartAPI(ctx, r, "8889", "test")
		time.Sleep(time.Second * 5)
		b, err := GetOk(8889)

		assert.Nil(t, err, "http get should return ok")
		assert.Equal(t, string(b), "ok", "response body should return ok")
		cancel()
		domain.WaitUntilAllTasksDone()
		b, err = GetOk(8889)
		assert.Nil(t, b)
		assert.NotNil(t, err, "http get should return error")
	})

}

func GetOk(port int) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/test", port))
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	return b, err
}
