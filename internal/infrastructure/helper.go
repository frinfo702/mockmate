package infrastructure

import (
	"log"
	"path/filepath"
	"strings"
)

const (
	fileLoadErrorMessage              = "ファイルの読み込みに失敗しました"
	yamlParseErrorMessage             = "yamlファイルの解析に失敗しました"
	notSupportedExtensionErrorMessage = "サポートされていない拡張子です"
)

// 拡張子は小文字に変換してから検証される。
// ok: .yaml, .yml, .YAML, .YML
func isYaml(fileName string) bool {
	ext := filepath.Ext(fileName)
	if strings.ToLower(ext) == ".yaml" || strings.ToLower(ext) == ".yml" {
		return true
	}
	log.Printf("%s: %s\n", notSupportedExtensionErrorMessage, ext)
	return false
}
