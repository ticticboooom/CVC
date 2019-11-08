package commands

import (
	"cvc/constants"
	"cvc/io"
	"flag"
	"fmt"
	"os"
	"strings"
)

func ParseRunInit(set *flag.FlagSet) {
	if checkAlreadyInRepository() {
		fmt.Print(constants.MessageAlreadyInRepository)
		os.Exit(1)
	}
	createCvcDirectory()
	dir, _ := os.Getwd()
	createRepositoryConfigFile(getRepositoryName(), dir)
}

func createCvcDirectory() {
	err := os.MkdirAll(fmt.Sprintf("./%s/", constants.DirectoryName), os.ModePerm)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func getRepositoryName() string {
	dir, _ := os.Getwd()
	baseDirSections := strings.Split(dir, string(byte(os.PathSeparator)))
	baseDir := baseDirSections[len(baseDirSections)-1]
	repo := strings.ReplaceAll(baseDir, " ", "-")
	return repo
}

func createRepositoryConfigFile(repoName string, rootDir string) {
	config := io.RepositoryConfig{
		Name: repoName,
	}
	WriteRepositoryConfig(config, rootDir)
	WriteFileListConfig(io.RepositoryFileList{Files: make([]string, 0)}, rootDir)
}

func checkAlreadyInRepository() bool {
	wd, _ := os.Getwd()

	_, err := io.FindRepositoryRoot(wd)
	if err != nil {
		return false
	}
	return true
}
