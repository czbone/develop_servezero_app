package config

import "github.com/go-sql-driver/mysql"

type Env struct {
	Database     mysql.Config
	MaxIdleConns int
	MaxOpenConns int
	ServerPort   string

	DatabaseName string

	RedisIp        string
	RedisPort      string
	RedisPassword  string
	RedisDb        int
	RedisSessionDb int
	RedisCacheDb   int

	SessionKey    string
	SessionSecret string

	AppSecret string

	// ログ機能
	AccessLog       bool
	AccessLogPath   string
	ErrorLog        bool
	ErrorLogPath    string
	SecurityLog     bool
	SecurityLogPath string
	DebugLog        bool
	DebugLogPath    string

	SqlLog bool

	TemplatePath string

	// サイト定義
	DefaultLanguage string
	Title           string
}

func GetEnv() *Env {
	return &env
}
