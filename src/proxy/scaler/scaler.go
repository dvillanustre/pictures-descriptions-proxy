package scaler

import (
	"bytes"
	"errors"
	"os/exec"
)

func GetPictureDefault(rawDescription []byte, version string) ([]byte, error) {
	cmd := exec.Command("/home/dvillanustre/Documents/apps/wkhtmltox/bin/wkhtmltoimage", "-", "-")
	cmd.Stdin = bytes.NewReader(rawDescription)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return nil, errors.New("Unexpected error when convert preview image")
	}

	return out.Bytes(), nil
}

func GetPictureWebp(rawDescription []byte, version string) ([]byte, error) {
	return nil, nil
}
