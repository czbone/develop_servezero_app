package config

var env = Env{
	// アプリケーション
	AppName:         "ServeZero",
	AppFilename:     "servezero",
	AppSecret:       "something-very-secret",
	DefaultLanguage: "ja",
	ServerPort:      "8080",

	// データベース
	DatabaseName: "zero.sqlite3",
	DatabasePath: "/usr/local/servezero/config",

	// ログ機能
	AccessLog:       true,
	AccessLogPath:   "log/access.log",
	ErrorLog:        true,
	ErrorLogPath:    "log/error.log",
	SecurityLog:     true,
	SecurityLogPath: "log/security.log",
	DebugLog:        true,
	DebugLogPath:    "log/debug.log",

	// テンプレートディレクトリ
	TemplatePath: "templates",

	// Nginx設定ファイル
	NginxUid:                 101,                                                  // Nginxプロセス実行ユーザ
	NginxSiteConfPath:        "/usr/local/servezero/volumes/nginx/sites-available", // サイト定義ファイル格納用
	NginxSiteConfTemplateDir: "conf-templates",
	NginxVirtualHostHome:     "/usr/local/servezero/volumes/vhost", // Webサイトホームディレクトリ
	// 設定ファイル定義用
	NginxContainerVirtualHostHome: "/var/www/vhost",

	// デバッグ用
	DebugOutputDir: "_dest",
}
