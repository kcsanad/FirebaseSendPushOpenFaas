package commons

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// ReadSecret reads a secret from /var/openfaas/secrets or from
// env-var 'secret_mount_path' if set.
func ReadSecret(key string) (string, error) {
	basePath := "/var/openfaas/secrets/"
	if len(os.Getenv("SECRET_MOUNT_PATH")) > 0 {
		basePath = os.Getenv("SECRET_MOUNT_PATH")
	}

	readPath := path.Join(basePath, key)
	secretBytes, readErr := ioutil.ReadFile(readPath)
	if readErr != nil {
		return "", fmt.Errorf("unable to read secret: %s, error: %s", readPath, readErr)
	}
	val := strings.TrimSpace(string(secretBytes))
	return val, nil
}
