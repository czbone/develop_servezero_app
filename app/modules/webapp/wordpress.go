package webapp

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"web/config"
	"web/modules/log"
)

type wordpressApp struct {
	Webapp
}

const (
	DOWNLOAD_URL       = "https://ja.wordpress.org/latest-ja.tar.gz"
	DOWNLOAD_FILE_HEAD = "download-"
)

func (wordpressApp *wordpressApp) Install(path string) bool {
	// テンポラリファイル作成
	file, err := os.CreateTemp("", config.GetEnv().AppFilename+"-"+DOWNLOAD_FILE_HEAD)
	if err != nil {
		log.Fatal(err)
	}

	// WordPressのソースアーカイブをダウンロード
	err = downloadFile(file.Name(), DOWNLOAD_URL)
	if err == nil {
		log.Info("WordPress installed!")
	}

	r, err := os.Open(file.Name())
	if err != nil {
		fmt.Println("error")
	}
	ExtractTarGz(r, path)
	return true
}
func (wordpressApp *wordpressApp) Backup(path string) bool {
	return true
}

func ExtractTarGz(gzipStream io.Reader, destDir string) {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(filepath.Join(destDir, header.Name), 0755); err != nil {
				log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			outFile, err := os.Create(filepath.Join(destDir, header.Name))
			if err != nil {
				log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
			outFile.Close()

		default:
			log.Fatalf(
				"ExtractTarGz: uknown type: %s in %s",
				header.Typeflag,
				header.Name)
		}

	}
}
