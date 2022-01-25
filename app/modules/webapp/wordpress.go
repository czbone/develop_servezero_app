package webapp

import (
	"io/ioutil"
	"web/modules/log"
)

type wordpressApp struct {
}

func NewWordpressApp() *wordpressApp {
	return &wordpressApp{}
}

func (wordpressApp *wordpressApp) Install(src string, dest string) bool {
	bytesRead, err := ioutil.ReadFile(src)
	if err != nil {
		log.Error(err)
		return false
	}

	err = ioutil.WriteFile(dest, bytesRead, 0644)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}
