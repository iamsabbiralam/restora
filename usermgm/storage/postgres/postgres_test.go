package postgres

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

var _testStorage *Storage

func TestMain(m *testing.M) {
	const dbConnEnv = "DATABASE_CONNECTION"
	ddlConnStr := os.Getenv(dbConnEnv)
	if ddlConnStr == "" {
		log.Printf("%s is not set, skipping", dbConnEnv)
		return
	}

	var teardown func()
	_testStorage, teardown = NewTestStorage(ddlConnStr, filepath.Join("..", "..", "migrations"))

	exitCode := m.Run()
	defer os.Exit(exitCode)
	if teardown != nil {
		defer teardown()
	}
}

func newTestStorage(tb testing.TB) *Storage {
	if testing.Short() {
		tb.Skip("skipping tests that use postgres on -short")
	}

	return _testStorage
}
