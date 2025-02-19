// cmd/mockmate/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/frinfo702/mockmate/internal/adapter/httpserver"
	"github.com/frinfo702/mockmate/internal/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 設定ファイルの存在チェック＆自動生成（初回起動用）
	err := infrastructure.EnsureDefaultConfig("config/default.yaml")
	if err != nil {
		log.Fatalf("Default config generation failed: %v", err)
	}

	// 設定ファイルをロード
	cfg, err := infrastructure.LoadConfig("config/default.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Echoの初期化
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// APIエンドポイントのルーティング登録
	httpserver.RegisterAPIRoutes(e, cfg)

	// ダッシュボード用ルートの登録（静的ファイル配信＋API）
	httpserver.RegisterDashboardRoutes(e)
	initDashboard(e) // dashboard.go に定義

	// サーバー起動設定
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	s := &http.Server{
		Addr:         addr,
		ReadTimeout:  parseDuration(cfg.Server.ReadTimeout, 30*time.Second),
		WriteTimeout: parseDuration(cfg.Server.WriteTimeout, 30*time.Second),
	}

	log.Printf("Starting server at %s", addr)
	if err := e.StartServer(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func parseDuration(s string, def time.Duration) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		return def
	}
	return d
}
