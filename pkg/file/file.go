package file

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func HomeDir() string {
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE") // windows
	}
	return home
}

func PrevProfileFile() (string, error) {
	home := HomeDir()
	if home == "" {
		return "", errors.New("HOME or USERPROFILE environment variable not set")
	}
	return filepath.Join(home, ".gcps"), nil
}

func WriteLastProfile(value string) error {
	path, err := PrevProfileFile()
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(value), 0600)
}

func ReadLastProfile() (string, error) {
	path, err := PrevProfileFile()
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		return "", nil
	}
	return string(b), err
}
