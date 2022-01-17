package config

import "github.com/go-sql-driver/mysql"

type Env struct {
	// アプリケーション
	AppName   string
	AppSecret string

	// サイト定義
	DefaultLanguage string
	Title           string

	Database     mysql.Config
	MaxIdleConns int
	MaxOpenConns int
	ServerPort   string

	DatabaseName string
	DatabasePath string

	RedisIp        string
	RedisPort      string
	RedisPassword  string
	RedisDb        int
	RedisSessionDb int
	RedisCacheDb   int

	SessionKey    string
	SessionSecret string

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

	// テンプレートディレクトリ
	TemplatePath string

	// Nginx設定ファイル
	NginxSiteConfPath        string
	NginxSiteConfTemplateDir string
	NginxSiteConfDomainHome  string
}

func GetEnv() *Env {
	return &env
}
