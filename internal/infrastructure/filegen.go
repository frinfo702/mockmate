package infrastructure

import (
	"os"
	"path/filepath"
)

// DefaultConfigContent is the default YAML content for config/default.yaml.
const DefaultConfigContent = `server:
  host: "0.0.0.0"
  port: 8080
  read_timeout: "30s"
  write_timeout: "30s"
endpoints:
  - name: users
    versions:
      - version: v1
        method: GET
        path: /api/v1/users
        response:
          status: 200
          headers:
            Content-Type: application/json
          body: |
            {
              "users": [
                {"id": 1, "name": "Alice"},
                {"id": 2, "name": "Bob"}
              ]
            }
      - version: v2
        method: GET
        path: /api/v2/users
        response:
          status: 200
          headers:
            Content-Type: application/json
          body: |
            {
              "users": [
                {"id": 1, "name": "Alice", "role": "admin"},
                {"id": 2, "name": "Bob", "role": "user"}
              ]
            }
`

// EnsureDefaultConfig checks if the config file exists; if not, it generates one with default content.
func EnsureDefaultConfig(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// Create the config directory if necessary.
		dir := getDir(filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		// Write default content.
		if err := os.WriteFile(filePath, []byte(DefaultConfigContent), 0644); err != nil {
			return err
		}
	}
	return nil
}

func getDir(filePath string) string {
	dir, _ := os.Stat(filePath)
	if dir == nil {
		// simple way: use filepath.Dir
		return filepath.Dir(filePath)
	}
	return ""
}
