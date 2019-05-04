package app

import (
	"os"
	"path/filepath"
	"strings"
)

func isExt(fileInfo os.FileInfo, ext string) bool {
	fileExt := filepath.Ext(fileInfo.Name())
	return strings.EqualFold(ext, fileExt)
}
