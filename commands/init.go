package commands

import (
	"cvc/constants"
	"cvc/io"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ParseRunInit(set *flag.FlagSet) {
	createCvcDirectory()
	dir, _ := os.Getwd()
	createRepositoryConfigFile(*getRepositoryName(), dir)

}

func createCvcDirectory() {
	if _, err := os.Stat(fmt.Sprintf("./%s/", constants.DirectoryName)); !os.IsNotExist(err) {
		fmt.Print("\nCould not init repository, it seems you already have one here silly.\n")
		os.Exit(1)
	}
	err := os.MkdirAll(fmt.Sprintf("./%s/", constants.DirectoryName), os.ModePerm)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func getRepositoryName() *string {
	dir, _ := os.Getwd()
	baseDirSections := strings.Split(dir, strconv.QuoteRune(os.PathSeparator))
	baseDir := baseDirSections[len(baseDirSections)-1]
	repo := flag.String("repo", strings.Replace(baseDir, " ", "-", 0), "The Name of your repository")
	return repo
}

func createRepositoryConfigFile(repoName string, rootDir string) {
	config := io.RepositoryConfig{
		Name: repoName,
	}
	WriteRepositoryConfig(config, rootDir)
	WriteFileListConfig(io.RepositoryFileList{Files: make([]string, 0)}, filepath.Join(rootDir, fmt.Sprintf("./%s/%s", constants.DirectoryName, constants.RepositoryConfigFileName)))
}
