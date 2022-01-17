package fileutil

import (
	"io/ioutil"
	"web/modules/log"
)

func CopyFile(src string, dest string) bool {
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
