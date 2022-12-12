package fs

import (
	"os"
	"path/filepath"
)

func WithAppRoot(path string) string {
	return filepath.Join(os.Getenv("APP_ROOT"), path)
}
