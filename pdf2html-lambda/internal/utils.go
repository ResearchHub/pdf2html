package internal

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"path"
)

func validateFileType(r io.Reader, contentType string) error {
	br := bufio.NewReaderSize(r, 512)

	// use the first 512 bytes so that we can check the content type
	// (peek does not advance the reader)
	b, err := br.Peek(512)
	if err != nil {
		return err
	}

	// ensure the upload file is a the given content type
	detectedContentType := http.DetectContentType(b)
	if detectedContentType != contentType {
		return fmt.Errorf("expected input to be [%s], but got [%s]",
			contentType, detectedContentType)
	}

	return nil
}

func replaceExtension(filepath string, newExt string) string {
	ext := path.Ext(filepath)
	idx := len(filepath) - len(ext)
	return filepath[0:idx] + newExt
}
