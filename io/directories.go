package io

import (
	"cvc/constants"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type FindRepositoryFailed struct{}

func (repo *FindRepositoryFailed) Error() string {
	return "Failed to find repository"
}

func FindRepositoryRoot(directory string) (string, error) {
	if _, err := os.Stat(filepath.Join(directory, fmt.Sprintf("./%s/", constants.DirectoryName))); !os.IsNotExist(err) {
		return directory, nil
	}
	baseDirSections := strings.Split(directory, strconv.QuoteRune(os.PathSeparator))
	if len(baseDirSections) < 2 {
		return "", &FindRepositoryFailed{}
	}
	baseDir := baseDirSections[0:(len(baseDirSections) - 1)]

	return FindRepositoryRoot(strings.Join(baseDir, strconv.QuoteRune(os.PathSeparator)))
}
