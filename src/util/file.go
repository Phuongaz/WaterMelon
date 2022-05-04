package util

import (
	"errors"
	"io/ioutil"
	"os"
)

var WorkingPath, _ = os.Getwd()

func MustReadFile(path string) []byte {
	return SelectAnyByteSlice(ioutil.ReadFile(path))
}

func MustDeleteFile(path string) {
	Must(os.Remove(path))
}

func MustWriteFile(path string, data []byte) {
	Must(os.WriteFile(path, data, 0666))
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func EnsureDir(dirName string) error {
	err := os.Mkdir(dirName, os.ModeDir)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}
