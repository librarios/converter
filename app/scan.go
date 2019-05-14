package app

import (
	"io/ioutil"
	"log"
	"sync"
)

type ScanCommand struct {
	bookInfoFileCh chan *BookInfoFile
}

type ScanOption struct {
	OutputFile string
}

func NewScanCommand() *ScanCommand {
	return &ScanCommand{
		bookInfoFileCh: make(chan *BookInfoFile, 100),
	}
}

func (c *ScanCommand) Scan(dir string, opt *ScanOption) error {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			bookInfoFile := <-c.bookInfoFileCh
			if bookInfoFile == nil {
				break
			}
			c.processBookInfoFile(bookInfoFile)
		}
	}()

	for _, fileInfo := range fileInfos {
		if isExt(fileInfo, ".txt") {
			bookInfoFile, err := parseBookInfoFile(dir, fileInfo, "")
			if err != nil {
				log.Printf("Failed to parse '%s': %s", fileInfo.Name(), err.Error())
			} else {
				c.bookInfoFileCh <- bookInfoFile
			}
		}
	}
	c.bookInfoFileCh <- nil
	wg.Wait()

	return nil
}

func (c *ScanCommand) processBookInfoFile(bookInfoFile *BookInfoFile) {
	log.Printf("READ: %s\n", bookInfoFile.fileInfo.Name())
	log.Printf("%v\n", bookInfoFile.bookInfo)
}
