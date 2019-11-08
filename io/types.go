package io

type RepositoryFileList struct {
	Files []string `yaml:"files"`
}

type RepositoryConfig struct {
	Name string `yaml:"name"`
}
