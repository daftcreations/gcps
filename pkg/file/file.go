package file

import (
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"path/filepath"
)

func PrevProfileFile() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".gcps"), nil
}

func WriteLastProfile(value string) error {
	path, err := PrevProfileFile()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
		return err
	}

	return ioutil.WriteFile(path, []byte(value), 0600)
}

func ReadLastProfile() (string, error) {
	path, err := PrevProfileFile()
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadFile(filepath.Clean(path))
	if os.IsNotExist(err) {
		return "", nil
	}

	return string(b), err
}
