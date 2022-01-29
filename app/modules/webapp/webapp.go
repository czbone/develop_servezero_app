package webapp

import (
	"fmt"
	"io"
	"mime"
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

func downloadFile(filepath string, url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	filenameCont := resp.Header.Get("Content-Disposition")
	mediaType, params, err := mime.ParseMediaType(filenameCont)
	if err != nil {
		fmt.Println("**Normal Filename error:", err)
	}
	filename := params["filename"]

	// mediaType = attachment
	fmt.Println("Normal:", mediaType, params)
	fmt.Print()

	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return filename, err
}
