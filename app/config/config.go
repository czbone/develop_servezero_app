package config

type Env struct {
	// アプリケーション
	AppName         string
	AppFilename     string
	AppSecret       string
	DefaultLanguage string
	ServerPort      string

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

	// テンプレートディレクトリ
	TemplatePath string

	// Nginx設定ファイル
	NginxUser                string // Nginxプロセス実行ユーザ
	NginxSiteConfPath        string
	NginxSiteConfTemplateDir string
	NginxVirtualHostHome     string // Webサイトホームディレクトリ
	// 設定ファイル定義用
	NginxContainerVirtualHostHome string

	// デバッグ用
	DebugOutputDir string
}

func GetEnv() *Env {
	return &env
}
