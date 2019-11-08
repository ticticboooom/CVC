package io

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return file
}

func WriteFile(contents []byte, path string) {
	err := ioutil.WriteFile(path, contents, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
