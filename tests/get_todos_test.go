//go:build integration

package tests

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTodos(t *testing.T) {
	spawnWebServerContainer(t)
	client := http.Client{}
	t.Run("Valid HTTP Request", func(t *testing.T) {
		r, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://127.0.0.1:8080/todos", nil)
		require.NoError(t, err)
		res, err := client.Do(r)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)
		baseUrl, err := url.Parse("http://127.0.0.1:80/render")
		require.NoError(t, err)
		params := url.Values{}
		params.Add("target", "webserver.get_todos.success")
		params.Add("from", "-5min")
		params.Add("format", "json")
		baseUrl.RawQuery = params.Encode()
		require.NoError(t, err)
		r, err = http.NewRequestWithContext(context.Background(), http.MethodGet, baseUrl.String(), nil)
		require.NoError(t, err)
		require.EventuallyWithT(t, func(collect *assert.CollectT) {
			isMetricEmitted, err := isMetricEmitted(client, r, "webserver.get_todos.success", 1)
			require.NoError(collect, err)
			require.True(collect, isMetricEmitted)
		}, time.Second*30, time.Second*3, "metric not emitted enough times")
	})
}
