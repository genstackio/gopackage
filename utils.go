package gopackage

import (
	"archive/zip"
	"bytes"
	"github.com/ohoareau/goaws/s3"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func newfileUploadRequest(uri string, params map[string]string, buf *bytes.Buffer) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	part, err := writer.CreateFormFile("file", "zip")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, buf)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func getObjectInByte(att *File) error {
	var err error
	cmp := strings.Split(att.Source, "/")
	if cmp[0] == "s3:" {
		att.Content, err = s3.GetObject(cmp[2], cmp[3]+"/"+cmp[4])
		if err != nil {
			return err
		}
		return nil
	}
	if cmp[0] == "http:" || cmp[0] == "https:" {
		file, err := http.Get(att.Source)
		if err != nil {
			return err
		}

		att.Content, err = io.ReadAll(file.Body)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func addToZip(w *zip.Writer, File File) error {
	f, err := w.Create(File.Name)
	if err != nil {
		return err
	}
	_, err = f.Write(File.Content)
	if err != nil {
		return err
	}
	return nil
}
