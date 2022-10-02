package model

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type callbackFn func()

func testGetTempDB(t *testing.T) (*gorm.DB, callbackFn) {
	dir, err := os.MkdirTemp("", "learne_model")
	require.NoError(t, err)

	db, err := gorm.Open(sqlite.Open(filepath.Join(dir, "test.db")), &gorm.Config{})
	require.NoError(t, err)
	return db, func() {
		os.RemoveAll(dir)
	}
}
