package server

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/chutommy/smart-passwd/pkg/config"
	"github.com/chutommy/smart-passwd/pkg/engine"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestNewServer(t *testing.T) { //nolint:tparallel
	t.Parallel()

	tests := []struct {
		name string
		cfg  *config.Config
		e    *engine.Engine
	}{
		{
			name: "ok",
			cfg: &config.Config{
				HTTPPort: 8080,
				Debug:    false,
				RootPath: "../..",
			},
			e: testEngine,
		},
		{
			name: "nil testEngine",
			cfg: &config.Config{
				HTTPPort: 8080,
				Debug:    false,
				RootPath: "../..",
			},
			e: nil,
		},
		{
			name: "invalid testEngine",
			cfg: &config.Config{
				HTTPPort: 8080,
				Debug:    false,
				RootPath: "../..",
			},
			e: &engine.Engine{},
		},
		{
			name: "with debug",
			cfg: &config.Config{
				HTTPPort: 8080,
				Debug:    true,
				RootPath: "../..",
			},
			e: testEngine,
		},
		{
			name: "port 80",
			cfg: &config.Config{
				HTTPPort: 80,
				Debug:    false,
				RootPath: "../..",
			},
			e: testEngine,
		},
		{
			name: "no config",
			cfg: &config.Config{
				Debug:    false,
				RootPath: "../..",
			},
			e: testEngine,
		},
	}

	for _, tt := range tests { //nolint:paralleltest
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			srv := NewServer(tt.cfg, tt.e)

			require.Equal(t, fmt.Sprint(":", tt.cfg.HTTPPort), srv.srv.Addr)
			require.Equal(t, tt.cfg.Debug, gin.Mode() == gin.DebugMode)
			require.Equal(t, tt.e, srv.engine)
		})
	}

	gin.SetMode(gin.TestMode)
}

func TestServer_Start(t *testing.T) {
	t.Parallel()

	s := NewServer(testConfig, testEngine)
	go require.NoError(t, s.srv.Close())
	require.True(t, errors.Is(s.Start(), http.ErrServerClosed))
}

func TestServer_Shutdown(t *testing.T) {
	t.Parallel()

	wg := &sync.WaitGroup{}

	srvOk := NewServer(testConfig, testEngine)

	wg.Add(1)

	go func() {
		require.NoError(t, srvOk.Shutdown(500*time.Millisecond))
		wg.Done()
	}()

	require.True(t, errors.Is(srvOk.Start(), http.ErrServerClosed))
	wg.Wait()

	srvErr := NewServer(testConfig, testEngine)

	wg.Add(1)

	go func() {
		require.Error(t, srvErr.Shutdown(0))
		wg.Done()
	}()

	require.True(t, errors.Is(srvErr.Start(), http.ErrServerClosed))
	wg.Wait()
}
