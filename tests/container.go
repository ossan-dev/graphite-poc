package tests

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

func spawnWebServerContainer(t *testing.T) {
	t.Helper()
	os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")
	compose, err := tc.NewDockerComposeWith(tc.WithStackFiles("../docker-compose.yml"))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal))
	})
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	err = compose.
		Up(ctx, tc.Wait(true))
	require.NoError(t, err)
}
