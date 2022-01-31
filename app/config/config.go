package config

type Env struct {
	// アプリケーション
	AppName         string
	AppFilename     string
	AppSecret       string
	DefaultLanguage string

	// システム環境
	ProductPath    string // 製品インストールディレクトリ
	ServerPort     string
	TemplatePath   string // テンプレートディレクトリ
	DebugOutputDir string // デバッグ用
	OnProductEnv   bool   // 製品環境で稼働しているかどうか

	// データベース
	DatabaseName string
	DatabasePath string

	// ログ機能
	AccessLog       bool
	AccessLogPath   string
	ErrorLog        bool
	ErrorLogPath    string
	SecurityLog     bool
	SecurityLogPath string
	DebugLog        bool
	DebugLogPath    string

	// Nginx設定ファイル
	NginxUser                string // Nginxプロセス実行ユーザ
	NginxSiteConfPath        string
	NginxSiteConfTemplateDir string
	NginxVirtualHostHome     string // Webサイトホームディレクトリ
	// 設定ファイル定義用
	NginxContainerVirtualHostHome string
}

func GetEnv() *Env {
	return &env
}
