package photoproc

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func CopyPhoto(uploadID string, file multipart.File, fileHeader *multipart.FileHeader) error {
	defer file.Close()

	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		return err
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		return err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		return err
	}

	dst, err := os.Create(fmt.Sprintf("./uploads/%s%s", uploadID, filepath.Ext(fileHeader.Filename)))
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	return nil

}
