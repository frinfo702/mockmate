package infrastructure

import "github.com/frinfo702/mockmate/internal/domain/entitiy"

type ConfigLoader interface {
	Load(filePath string) (*entitiy.Config, error)
}
