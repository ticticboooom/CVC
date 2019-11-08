package io

import (
	"cvc/constants"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func FindRepositoryRoot(directory string) string {
	if _, err := os.Stat(filepath.Join(directory, fmt.Sprintf("./%s/", constants.DirectoryName))); !os.IsNotExist(err) {
		return directory
	}
	baseDirSections := strings.Split(directory, strconv.QuoteRune(os.PathSeparator))
	baseDir := baseDirSections[0:(len(baseDirSections) - 1)]

	return FindRepositoryRoot(strings.Join(baseDir, strconv.QuoteRune(os.PathSeparator)))
}
