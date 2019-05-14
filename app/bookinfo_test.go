package app

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestParseBookInfoFile(t *testing.T) {
	a := assert.New(t)

	baseDir := "../test/info-files/"
	filenames := map[string]string{
		"euc-kr.txt": "euc-kr",
		"utf-8.txt": "utf-8",
	}

	for filename, enc := range filenames {
		path := filepath.Join(baseDir, filename)
		fileInfo, err := os.Stat(path)
		if err != nil {
			t.Error(err)
		}
		bookInfoFile, err := parseBookInfoFile(baseDir, fileInfo, enc)
		if err != nil {
			t.Error(err)
		}

		bookInfo := bookInfoFile.bookInfo
		a.Equal(baseDir, bookInfoFile.dir, "dir mismatch")
		a.Equal("9788971723760", bookInfo.Isbn)
		a.Equal("홍길동", bookInfo.Owner)
		a.Equal("1909", bookInfo.OrigPubDate)
		a.Equal("The Great Gatsby", bookInfo.OrigTitle)
		a.Equal("2005-02-25", bookInfo.PubDate)
		a.Equal("2010-03-18", bookInfo.BoughtDate)
		a.Equal("1900", bookInfo.BoughtPrice)
		a.Equal("2011.01.27", bookInfo.ScanDate)
		a.Equal(144, bookInfo.ScanPages)
		a.Equal("5000", bookInfo.Price)
	}
}
