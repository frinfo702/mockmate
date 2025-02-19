// このファイルは、ダッシュボード専用のルート設定および静的ファイル配信用です。
package main

import (
	"github.com/labstack/echo/v4"
)

func initDashboard(e *echo.Echo) {
	// /dashboard ルートに対して、dashboardディレクトリ内の静的ファイルを配信
	e.Static("/dashboard", "dashboard")
}
