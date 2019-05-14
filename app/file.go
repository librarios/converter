package app

import (
	"github.com/thoas/go-funk"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// isExt checks if extension matches
func isExt(fileInfo os.FileInfo, ext string) bool {
	fileExt := filepath.Ext(fileInfo.Name())
	return strings.EqualFold(ext, fileExt)
}

// listDir list all files with given extension
func listDir(dir string, ext string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return funk.Filter(files, func(file os.FileInfo) bool {
		return isExt(file, ext)
	}).([]os.FileInfo), nil
}
