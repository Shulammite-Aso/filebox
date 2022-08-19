package util

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err := os.WriteFile(fileName, body, 0666); err != nil {
		return err
	}
	return nil
}

func VerifySize(name string) error {
	const MAX_File_SIZE = 1024 * 1024 // 1MB

	stats, err := os.Stat(name)
	if err != nil {
		return err
	}

	fmt.Printf("Size: %d\n", stats.Size())

	fileSize := stats.Size()

	if fileSize > MAX_File_SIZE {
		return errors.New("too big: file must be 1MB or less")
	}

	return nil

}

func HandleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
