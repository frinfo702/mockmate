package infrastructure

import (
	"os"
	"testing"
)

// TestFileConfigLoader_Load_Success
// ・一時ファイルに有効なYAMLを出力し、Loadメソッドでパースできるかを検証します。
// ・想定される ServerConfig のフィールドとして Host と Port をチェックしています。
func TestFileConfigLoader_Load_Success(t *testing.T) {
	// 一時ファイルを作成し、YAMLコンテンツを書き込む
	content := `
host: "localhost"
port: 8080
`
	tmpfile, err := os.CreateTemp("", "testconfig-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpfile.Close()

	loader := NewFileConfigLoader()
	config, err := loader.Load(tmpfile.Name())
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if config == nil {
		t.Fatalf("expected non-nil config")
	}

	// ※以下は、entitiy.ServerConfigに host, port のフィールドが存在する前提です
	if config.Host != "localhost" {
		t.Errorf("expected host 'localhost', got '%s'", config.Host)
	}
	if config.Port != 8080 {
		t.Errorf("expected port 8080, got %d", config.Port)
	}
}

// TestFileConfigLoader_Load_FileNotFound
// ・存在しないファイルを指定した場合、エラーが返されることを検証します。
func TestFileConfigLoader_Load_FileNotFound(t *testing.T) {
	nonExistentFile := "nonexistent.yaml"
	loader := NewFileConfigLoader()
	_, err := loader.Load(nonExistentFile)
	if err == nil {
		t.Fatalf("expected error for non-existent file, got nil")
	}
}

// TestFileConfigLoader_Load_InvalidYAML
// ・不正なYAML形式のファイルを指定した場合、アンマーシャルエラーが返されることを検証します。
func TestFileConfigLoader_Load_InvalidYAML(t *testing.T) {
	// 一時ファイルに不正なYAMLを出力
	content := `
invalid_yaml: [this, is, not: valid: yaml
`
	tmpfile, err := os.CreateTemp("", "invalidconfig-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpfile.Close()

	loader := NewFileConfigLoader()
	_, err = loader.Load(tmpfile.Name())
	if err == nil {
		t.Fatalf("expected error for invalid YAML content, got nil")
	}
}
