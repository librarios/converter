package app

import (
	"bytes"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"path/filepath"
)
import "gopkg.in/ini.v1"

type BookInfo struct {
	Isbn        string
	Owner       string
	OrigPubDate string
	OrigTitle   string
	PubDate     string
	ScanDate    string
	ScanPages   int
	BoughtDate  string
	BoughtPrice string
	Price       string
}

type BookInfoFile struct {
	bookInfo *BookInfo
	dir      string
	fileInfo os.FileInfo
}

func NewBookInfo() *BookInfo {
	return &BookInfo{}
}

func parseBookInfoFile(dir string, fileInfo os.FileInfo, enc string) (*BookInfoFile, error) {
	path := filepath.Join(dir, fileInfo.Name())
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	bookInfo, err := parseBookInfoBytes(bytes, enc)
	if err != nil {
		return nil, err
	}

	return &BookInfoFile{
		bookInfo: bookInfo,
		dir:      dir,
		fileInfo: fileInfo,
	}, nil
}

func decodeEuckr(buf []byte) ([]byte, error) {
	var bufs bytes.Buffer
	wr := transform.NewWriter(&bufs, korean.EUCKR.NewDecoder())
	defer wr.Close()
	if _, err := wr.Write(buf); err != nil {
		return nil, err
	}
	return bufs.Bytes(), nil
}

func parseBookInfoBytes(buf []byte, enc string) (*BookInfo, error) {
	if enc == "euc-kr" {
		if euckrBuf, err := decodeEuckr(buf); err != nil {
			return nil, err
		} else {
			buf = euckrBuf
		}
	}

	cfg, err := ini.Load(buf)
	if err != nil {
		return nil, err
	}

	section := cfg.Section("")
	bookInfo := NewBookInfo()
	bookInfo.Isbn = section.Key("isbn").String()
	bookInfo.Owner = section.Key("owner").String()
	bookInfo.OrigPubDate = section.Key("origPubDate").String()
	bookInfo.OrigTitle = section.Key("origTitle").String()
	bookInfo.PubDate = section.Key("pubDate").String()
	bookInfo.ScanDate = section.Key("scanDate").String()
	if bookInfo.ScanPages, err = section.Key("scanPages").Int(); err != nil {
		bookInfo.ScanPages = 0
	}
	bookInfo.BoughtDate = section.Key("boughtDate").String()
	bookInfo.BoughtPrice = section.Key("boughtPrice").String()
	bookInfo.Price = section.Key("price").String()

	return bookInfo, nil
}
