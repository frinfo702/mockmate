package httpserver

import (
	"net/http"

	"github.com/frinfo702/mockmate/internal/entity"
	"github.com/frinfo702/mockmate/internal/usecase"
	"github.com/labstack/echo/v4"
)

// RegisterAPIRoutes registers API routes based on the loaded config.
func RegisterAPIRoutes(e *echo.Echo, cfg *entity.Config) {
	// 各エンドポイントごとにルートを登録する
	for _, ep := range cfg.EndPoints {
		for _, ver := range ep.Versions {
			// Capture variables for closure
			method := ver.Method
			path := ver.Path
			response := ver.Response

			handler := func(c echo.Context) error {
				// クエリパラメータでバージョン指定があれば、ユースケース層で選択
				queryVer := c.QueryParam("version")
				selected := usecase.SelectVersion(ver, queryVer)
				// レスポンスヘッダーの設定
				for k, v := range selected.Response.Header {
					c.Response().Header().Set(k, v)
				}
				return c.String(selected.Response.Status, selected.Response.Body)
			}

			// Echoのルーティングに登録
			switch method {
			case http.MethodGet:
				e.GET(path, handler)
			case http.MethodPost:
				e.POST(path, handler)
			case http.MethodPut:
				e.PUT(path, handler)
			case http.MethodDelete:
				e.DELETE(path, handler)
			default:
				e.Add(method, path, handler)
			}
		}
	}
}
