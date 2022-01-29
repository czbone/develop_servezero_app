package webapp

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"web/modules/log"
)

// アーカイブファイルの解凍
// srcFile: ソースファイルパス
// filename: HTTPレスポンスヘッダから取得した正式なファイル名
// destDir: 解凍先ディレクトリ
// 返り値: 解凍後生成されたディレクトリ, エラーステータス
func extract(srcFile string, filename string, destDir string) (string, error) {
	var err error
	var extractedDir string

	// 拡張子取得
	ext := getExt(filename)
	fmt.Println(srcFile + " - " + ext)

	switch ext {
	case ".tar.gz":
		fmt.Println("archiving..." + ext)
		extractedDir, err = extractTarGz(srcFile, destDir)
		if err == nil {
			fmt.Println("archived: " + extractedDir)
		}
	}
	return extractedDir, err
}

func getExt(filename string) string {
	base := filepath.Base(filename)
	ext := strings.ToLower(filepath.Ext(base))
	nonExt := filename[:len(filename)-len(ext)]
	if strings.ToLower(filepath.Ext(nonExt)) == ".tar" {
		ext = ".tar" + ext
	}
	return ext
}

func extractTarGz(srcFile string, destDir string) (string, error) {
	gzipStream, err := os.Open(srcFile)
	if err != nil {
		log.Error("extractTarGz: Open() failed: %s", err.Error())
		return "", err
	}
	defer gzipStream.Close()

	// gzip解凍
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		log.Error("extractTarGz: NewReader() failed: %s", err.Error())
		return "", err
	}

	var isFirst = true
	var extractedDir = ""
	tarReader := tar.NewReader(uncompressedStream)
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Error("extractTarGz: Next() failed: %s", err.Error())
			return "", err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(filepath.Join(destDir, header.Name), 0755); err != nil {
				log.Error("extractTarGz: Mkdir() failed: %s", err.Error())
				return "", err
			}
			if isFirst {
				extractedDir = filepath.Join(destDir, header.Name)
			}
		case tar.TypeReg:
			if isFirst {
				// ルートがディレクトリでない場合はディレクトリを作成
				addDirName := strings.TrimSuffix(filepath.Base(srcFile), ".tar.gz")
				destDir = filepath.Join(destDir, addDirName) // 解凍先ルートディレクトリを変更

				// ルートディレクトリを更新
				if err := os.Mkdir(destDir, 0755); err != nil {
					log.Error("extractTarGz: Mkdir() failed: %s", err.Error())
					return "", err
				}
				extractedDir = destDir
			}
			outFile, err := os.Create(filepath.Join(destDir, header.Name))
			if err != nil {
				log.Error("extractTarGz: Create() failed: %s", err.Error())
				return "", err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Error("extractTarGz: Copy() failed: %s", err.Error())
				return "", err
			}
			outFile.Close()

		default:
			log.Error("extractTarGz: uknown type: %d in %s", header.Typeflag, header.Name)
			return "", err
		}

		isFirst = false
	}
	return extractedDir, nil
}
