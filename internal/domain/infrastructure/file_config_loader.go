package infrastructure

import (
	"log"
	"os"

	"github.com/frinfo702/mockmate/internal/domain/entitiy"
	"gopkg.in/yaml.v2"
)

type FileConfigLoader struct {
}

func NewFileConfigLoader() *FileConfigLoader {
	return &FileConfigLoader{}
}

// ファイルパスでyamlファイルを指定する。
// yamlファイルをパースしてモデルに代入して返す.
// yamlファイルの拡張子はYAMLとかでも動いてしまうので変換する。
func (f *FileConfigLoader) Load(filePath string) (*entitiy.ServerConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(fileLoadErrorMessage, err)
		return &entitiy.ServerConfig{}, err
	}

	var config entitiy.ServerConfig
	if isYaml(filePath) {
		err := yaml.Unmarshal(data, &config)
		if err != nil {
			log.Println(yamlParseErrorMessage, err)
			return &entitiy.ServerConfig{}, err
		}
	}

	return &config, nil
}
