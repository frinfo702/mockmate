package infrastructure

import "github.com/frinfo702/mockmate/internal/entity"

type ConfigLoader interface {
	LoadConfig(filePath string) (*entity.Config, error)
}
