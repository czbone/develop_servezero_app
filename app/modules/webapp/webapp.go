package webapp

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	WordPressWebAppType = "wordpress"
)

type Webapp interface {
	Install(path string) bool
	Backup(path string) bool
}

func NewWebapp(appType string) (Webapp, error) {
	// Webアプリケーションに合わせてインスタンス作成
	switch appType {
	case WordPressWebAppType:
		return &wordpressApp{}, nil
	}
	return nil, fmt.Errorf("wrong webapp type")
}

func downloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
