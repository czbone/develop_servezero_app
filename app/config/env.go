package config

import "github.com/go-sql-driver/mysql"

var env = Env{
	// アプリケーション
	AppName:   "ServeZero",
	AppSecret: "something-very-secret",

	// サイト定義
	DefaultLanguage: "ja",
	Title:           "ServeZero",

	ServerPort: "8080",

	Database: mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.0.1:3306",
		DBName:               "gin-template",
		Collation:            "utf8mb4_unicode_ci",
		Net:                  "tcp",
		AllowNativePasswords: true,
	},
	MaxIdleConns: 50,
	MaxOpenConns: 100,
	DatabaseName: "zero.sqlite3",
	DatabasePath: "/usr/local/servezero/config",

	RedisIp:       "127.0.0.1",
	RedisPort:     "6379",
	RedisPassword: "",
	RedisDb:       0,

	RedisSessionDb: 1,
	RedisCacheDb:   2,

	// ログ機能
	AccessLog:     true,
	AccessLogPath: "log/access.log",

	ErrorLog:     true,
	ErrorLogPath: "log/error.log",

	SecurityLog:     true,
	SecurityLogPath: "log/security.log",

	DebugLog:     true,
	DebugLogPath: "log/debug.log",

	// テンプレートディレクトリ
	TemplatePath: "templates",

	// Nginx設定ファイル
	NginxSiteConfPath:        "/usr/local/servezero/volumes/nginx/sites-available",
	NginxSiteConfTemplateDir: "conf-templates",
	NginxSiteConfDomainHome:  "/usr/local/servezero/volumes/nginx/home",
}
