package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// add cleanup of containers
	exitCode := m.Run()
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}
