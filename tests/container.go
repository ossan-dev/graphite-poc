package tests

import (
	"context"
	"os"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/errdefs"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

func reRunContainersAfterConflict(ctx context.Context, composeReq tc.ComposeStack) error {
	dockerClient, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}
	dockerClient.NegotiateAPIVersion(ctx)
	containers := composeReq.Services()
	for _, c := range containers {
		err = dockerClient.ContainerStop(ctx, c, container.StopOptions{})
		if err != nil {
			return err
		}
		err = dockerClient.ContainerRemove(ctx, c, container.RemoveOptions{})
		if err != nil {
			return err
		}
	}
	return composeReq.
		Up(ctx, tc.Wait(true))
}

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
	if err != nil && errdefs.IsConflict(err) {
		require.NoError(t, reRunContainersAfterConflict(ctx, compose))
		return
	}
	require.NoError(t, err)
}
