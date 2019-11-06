package commands

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ParseRunInit() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
		return
	}
	baseDir := filepath.Base(dir)
	repo := flag.String("repo", strings.Replace(baseDir, " ", "-", 0), "The Name of your repository")

	createCvcDirectory()

}


func createCvcDirectory() {
	dir, err := os.Getwd()
	err = os.MkdirAll(dir+".cvc", os.ModePerm)
	if err != nil {
		fmt.Print(err)
		return
	}
}


