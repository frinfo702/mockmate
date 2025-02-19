package infrastructure

import (
	"log"
	"os"

	"github.com/frinfo702/mockmate/internal/entity"
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
func (f *FileConfigLoader) LoadConfig(filePath string) (*entity.ServerConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(fileLoadErrorMessage, err)
		return &entity.ServerConfig{}, err
	}

	var config entity.ServerConfig
	if isYaml(filePath) {
		err := yaml.Unmarshal(data, &config)
		if err != nil {
			log.Println(yamlParseErrorMessage, err)
			return &entity.ServerConfig{}, err
		}
	}

	return &config, nil
}
