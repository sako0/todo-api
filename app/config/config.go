package config

import (
	"fmt"
	"os"
)

type appConfig struct {
	AppInfo *AppInfo
}

type AppInfo struct {
	DatabaseURL string
}

func LoadConfig() (*appConfig, error) {
	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		return nil, fmt.Errorf("環境変数MYSQL_HOSTが設定されていません")
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		return nil, fmt.Errorf("環境変数MYSQL_USERが設定されていません")
	}
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		return nil, fmt.Errorf("環境変数MYSQL_PASSWORDが設定されていません")
	}
	mysqlDBName := os.Getenv("MYSQL_DATABASE")
	if mysqlDBName == "" {
		return nil, fmt.Errorf("環境変数MYSQL_DATABASEが設定されていません")
	}
	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlPort == "" {
		return nil, fmt.Errorf("環境変数MYSQL_PORTが設定されていません")
	}

	databaseURL := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDBName + "?charset=utf8mb4&parseTime=true"

	appInfo := &AppInfo{
		DatabaseURL: databaseURL,
	}

	config := appConfig{
		AppInfo: appInfo,
	}

	return &config, nil
}
