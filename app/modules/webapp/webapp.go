package webapp

import "fmt"

const (
	WordPressWebAppType = "wordpress"
)

type Webapp interface {
	Install() bool
	Backup() bool
}

func NewWebapp(appType string) (Webapp, error) {

	switch appType {
	case WordPressWebAppType:
		//return NewWordpressApp(), nil
		return &wordpressApp{}, nil
	}
	return nil, fmt.Errorf("wrong webapp type")
}
