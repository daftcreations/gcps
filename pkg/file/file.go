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

func WriteLastProfile(path, value string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(value), 0644)
}

func ReadLastProfile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		return "", nil
	}
	return string(b), err
}
