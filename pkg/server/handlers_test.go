package server

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chutommy/smart-passwd/pkg/engine"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestHomePageHandler(t *testing.T) {
	t.Parallel()

	r := gin.New()
	setRouter("../..", nil, r)

	w := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	require.NoError(t, err)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/html; charset=utf-8", w.Header().Get("content-type"))

	require.NotNil(t, w.Body)
	_, err = html.Parse(w.Body)
	require.NoError(t, err)
}

func TestPingHandler(t *testing.T) {
	t.Parallel()

	r := gin.New()
	setRouter("../..", nil, r)

	w := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/ping", nil)
	require.NoError(t, err)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, `{"status":"online"}`, w.Body.String())
}

func TestPasswordGenHandler(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		engine  *engine.Engine
		reqBody io.Reader
		expCode int
	}{
		{
			name:    "ok",
			engine:  testEngine,
			reqBody: strings.NewReader(`{"len": 20, "extra": 5, "helper": ""}`),
			expCode: 200,
		},
		{
			name:    "invalid password length",
			engine:  testEngine,
			reqBody: strings.NewReader(`{"len": 100, "extra": 0, "helper": ""}`),
			expCode: 400,
		},
		{
			name:    "invalid request params",
			engine:  testEngine,
			reqBody: strings.NewReader(`{"len": 10, "extra": 0, "helper": "foo"}`),
			expCode: 400,
		},
		{
			name:    "invalid request body",
			engine:  testEngine,
			reqBody: strings.NewReader(`{len: 10, extra: 0, helper: ""}`),
			expCode: 400,
		},
		{
			name:    "invalid engine",
			engine:  invalidEngine,
			reqBody: strings.NewReader(`{"len": 20, "extra": 5, "helper": ""}`),
			expCode: 500,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := gin.New()
			setRouter("../..", tt.engine, r)

			w := httptest.NewRecorder()
			req, err := http.NewRequestWithContext(context.Background(), "POST", "/gen", tt.reqBody)
			require.NoError(t, err)
			r.ServeHTTP(w, req)

			require.Equal(t, tt.expCode, w.Code)
			require.Equal(t, "application/json; charset=utf-8", w.Header().Get("content-type"))

			require.NotNil(t, w.Body)
			b, err := ioutil.ReadAll(w.Body)
			require.NoError(t, err)
			require.NoError(t, json.Unmarshal(b, &struct{}{}))
		})
	}
}
