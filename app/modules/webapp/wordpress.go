package webapp

import (
	"os"
	"web/config"
	"web/modules/log"
)

type wordpressApp struct {
	Webapp
}

const (
	DOWNLOAD_URL       = "https://ja.wordpress.org/latest-ja.tar.gz"
	DOWNLOAD_FILE_HEAD = "download-"
	DOWNLOAD_DIR_HEAD  = "download-dir-"
)

// Webアプリケーションのソースパッケージをダウンロードし、指定のパスに展開
// path=解凍したディレクトリの配置パス
func (wordpressApp *wordpressApp) Install(path string) bool {
	// テンポラリファイル作成
	tempFile, err := os.CreateTemp("", config.GetEnv().AppFilename+"-"+DOWNLOAD_FILE_HEAD)
	if err != nil {
		log.Error(err)
		return false
	}

	// WordPressのソースアーカイブをダウンロード
	filename, err := downloadFile(tempFile.Name(), DOWNLOAD_URL)
	if err != nil {
		log.Error(err)
		return false
	}

	// 解凍用の一時ディレクトリ作成
	destDir, err := os.MkdirTemp("", config.GetEnv().AppFilename+"-"+DOWNLOAD_DIR_HEAD)
	if err != nil {
		log.Error(err)
		return false
	}
	defer os.RemoveAll(destDir)

	// ソースを解凍
	extractedDir, err := extract(tempFile.Name(), filename, destDir)
	if err != nil {
		log.Error(err)
		return false
	}

	// 指定の位置にディレクトリを移動
	err = os.Rename(extractedDir, path)
	if err == nil {
		return true
	} else {
		log.Error(err)
		return false
	}
}
func (wordpressApp *wordpressApp) Backup(path string) bool {
	return true
}
