package httpserver

import (
	"net/http"

	"github.com/frinfo702/mockmate/internal/entity"
	"github.com/frinfo702/mockmate/internal/infrastructure"
	"github.com/labstack/echo/v4"
)

// RegisterDashboardRoutes registers routes for the dashboard API.
func RegisterDashboardRoutes(e *echo.Echo) {
	e.GET("/dashboard/config", func(c echo.Context) error {
		// ここでは、シンプルに設定ファイル（default.yaml）の内容を返す例
		// 実際には、内部で保持している設定情報を返す設計にする
		cfg, err := loadConfigForDashboard("config/default.yaml")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to load config"})
		}
		return c.JSON(http.StatusOK, cfg)
	})
}

func loadConfigForDashboard(path string) (*entity.Config, error) {
	// infrastructure.LoadConfigを利用
	return infrastructure.LoadConfig(path)
}
