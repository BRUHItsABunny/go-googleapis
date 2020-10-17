package go_googleapis

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
)

func (ij *ImageJob) GetImageBase64() error {
	var (
		err      error
		toEncode []byte
	)

	if len(ij.Bytes) != 0 {
		toEncode = ij.Bytes
	} else if (ij.Reader) != nil {
		toEncode, err = ioutil.ReadAll(ij.Reader)
	} else {
		err = errors.New("no source")
	}

	if err == nil {
		ij.Content = base64.URLEncoding.EncodeToString(toEncode)
	}

	return err
}
