package slsa

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// Sha256 returns a sha256 hex digest of the file contents at the given path.
func Sha256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", nil
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
