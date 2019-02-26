package main

import (
	"io/ioutil"
	"path/filepath"
)

// Explorer is
type Explorer struct{}

// ExploreFile return []string
func (e Explorer) ExploreFile(dir string, recursive bool) []string {
	var paths []string

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return paths
	}

	for _, file := range files {
		if !file.IsDir() {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}

		if file.IsDir() {
			if recursive {
				var dirpaths = e.ExploreFile(filepath.Join(dir, file.Name()), recursive)
				paths = append(paths, dirpaths...)
			}
		}
	}

	return paths
}

// ExploreDirectory return []string
func (e Explorer) ExploreDirectory(dir string, recursive bool) []string {
	var paths []string

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return paths
	}

	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, filepath.Join(dir, file.Name()))
			if recursive {
				var dirpaths = e.ExploreDirectory(filepath.Join(dir, file.Name()), recursive)
				paths = append(paths, dirpaths...)
			}
		}
	}

	return paths
}
