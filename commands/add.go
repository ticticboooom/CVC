package commands

import (
	"cvc/constants"
	"cvc/io"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func ParseRunAdd(set *flag.FlagSet) {
	showWorkingRepository()
	addFilesToRepositoryFromDirectory()
}

func addFilesToRepositoryFromDirectory() {
	if len(os.Args) < 3 {
		fmt.Print("\nYou must provide a relative path for the files you wish to add\n")
		os.Exit(1)
	}
	relativeDir := os.Args[2]
	absoluteDir, _ := filepath.Abs(relativeDir)
	wd, _ := os.Getwd()
	files := getFilesConfig(wd).Files

	err := filepath.Walk(absoluteDir, func(path string, info os.FileInfo, err error) error {
		if !StringArrayContains(files, path) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file + " +")
	}
	WriteFileListConfig(io.RepositoryFileList{Files: files}, wd)
}

func getFilesConfig(dir string) io.RepositoryFileList {
	if _, err := os.Stat(filepath.Join(io.FindRepositoryRoot(dir), fmt.Sprintf("./%s/", constants.DirectoryName), constants.RepositoryIncludedFilesFileName)); os.IsNotExist(err) {
		return io.RepositoryFileList{Files: make([]string, 0)}
	}
	fileConf := io.RepositoryFileList{}
	confContent := io.ReadFile(filepath.Join(io.FindRepositoryRoot(dir), fmt.Sprintf("./%s/", constants.DirectoryName), constants.RepositoryIncludedFilesFileName))
	err := yaml.Unmarshal(confContent, &fileConf)
	if err != nil {
		panic(err)
	}
	return fileConf
}
