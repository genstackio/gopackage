package gopackage

import (
	"archive/zip"
	"bytes"
	"net/http"
)

//goland:noinspection GoUnusedExportedFunction
func CreatePackage(packs *Package) (string, error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	for i := 0; i < len(packs.Files); i++ {
		err := getObjectInByte(&packs.Files[i])
		if err != nil {
			return "", err
		}
		err = addToZip(w, packs.Files[i])
		if err != nil {
			return "", err
		}
	}

	err := w.Close()
	if err != nil {
		return "", err
	}

	request, err := newfileUploadRequest(packs.Target.Location, packs.Target.Params, buf)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return "", err
	}

	return "", nil
}
