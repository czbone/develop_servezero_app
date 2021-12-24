package config

import "github.com/go-sql-driver/mysql"

var env = Env{
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
	DatabaseName: "zero.db",

	RedisIp:       "127.0.0.1",
	RedisPort:     "6379",
	RedisPassword: "",
	RedisDb:       0,

	RedisSessionDb: 1,
	RedisCacheDb:   2,

	// ログ機能
	AccessLog:     true,
	AccessLogPath: "storage/logs/access.log",

	ErrorLog:     true,
	ErrorLogPath: "storage/logs/error.log",

	SecurityLog:     true,
	SecurityLogPath: "storage/logs/security.log",

	DebugLog:     true,
	DebugLogPath: "storage/logs/debug.log",

	TemplatePath: "templates",

	AppSecret: "something-very-secret",

	// サイト定義
	DefaultLanguage: "ja",
	Title:           "ベースWebプログラム",
}
