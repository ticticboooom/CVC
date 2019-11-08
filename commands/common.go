package commands

import (
	"cvc/constants"
	"cvc/io"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func showWorkingRepository() {
	dir, _ := os.Getwd()
	repositoryRoot, err := io.FindRepositoryRoot(dir)
	if err != nil {
		fmt.Print(constants.MessageNotInRepository)
		os.Exit(1)
	}
	config := io.RepositoryConfig{}
	confContent := io.ReadFile(filepath.Join(repositoryRoot, fmt.Sprintf("./%s/%s", constants.DirectoryName, constants.RepositoryConfigFileName)))
	err = yaml.Unmarshal(confContent, &config)
	if err != nil {
		panic(err)
	}
	fmt.Print("\nYou are in repository named: " + config.Name + "\n")
}

func StringArrayContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func WriteFileListConfig(config io.RepositoryFileList, wd string) {
	configToWrite, _ := yaml.Marshal(config)
	dir, err := io.FindRepositoryRoot(wd)
	if err != nil {
		fmt.Print(constants.MessageNotInRepository)
		os.Exit(1)
	}
	io.WriteFile(configToWrite, filepath.Join(dir, fmt.Sprintf("./%s/", constants.DirectoryName), constants.RepositoryIncludedFilesFileName))
}

func WriteRepositoryConfig(config io.RepositoryConfig, wd string) {
	content, _ := yaml.Marshal(config)
	io.WriteFile(content, filepath.Join(wd, fmt.Sprintf("./%s/%s", constants.DirectoryName, constants.RepositoryConfigFileName)))
}
