package apimanager_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/lawmatsuyama/transactions/domain"
	"github.com/lawmatsuyama/transactions/infra/apimanager"

	"github.com/stretchr/testify/assert"
)

func TestApiManager(t *testing.T) {
	t.Run("01.api manager should start and shutdown ok", func(t *testing.T) {
		defer t.Cleanup(func() {
			domain.CleanupWaitGroup()
		})
		r := chi.NewRouter()
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

		apimanager.StartAPI(context.TODO(), r, "8888", "test")
		b, err := GetOk()

		assert.Nil(t, err, "http get should return ok")
		assert.Equal(t, string(b), "ok", "response body should return ok")

		apimanager.ShutdownAPI()
		b, err = GetOk()
		assert.Nil(t, b)
		assert.NotNil(t, err, "http get should return error")
	})

	t.Run("02.api manager should start and graceful shutdown ok", func(t *testing.T) {
		defer t.Cleanup(func() {
			domain.CleanupWaitGroup()
		})
		r := chi.NewRouter()
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

		ctx, cancel := context.WithCancel(context.TODO())
		apimanager.StartAPI(ctx, r, "8888", "test")
		b, err := GetOk()

		assert.Nil(t, err, "http get should return ok")
		assert.Equal(t, string(b), "ok", "response body should return ok")
		cancel()
		domain.WaitUntilAllTasksDone()
		b, err = GetOk()
		assert.Nil(t, b)
		assert.NotNil(t, err, "http get should return error")
	})

}

func GetOk() ([]byte, error) {
	resp, err := http.Get("http://localhost:8888/test")
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	return b, err
}
