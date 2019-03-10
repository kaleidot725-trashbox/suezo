package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// Explorer is
type Explorer struct{}

// ExploreFile return []string
func (e Explorer) ExploreFile(dir string, recursive bool) (paths []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return paths, errors.New("not found directory")
	}

	for _, file := range files {
		if file.IsDir() {
			if recursive {
				var recursivePath = filepath.Join(dir, file.Name())
				var dirpaths, err = e.ExploreFile(filepath.Join(dir, file.Name()), recursive)
				if err != nil {
					println("warning not found directory " + recursivePath)
					continue
				}

				paths = append(paths, dirpaths...)
			}
		} else {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths, nil
}

// ExploreDirectory return []string
func (e Explorer) ExploreDirectory(dir string, recursive bool) (paths []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return paths, errors.New("not found directory")
	}

	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, filepath.Join(dir, file.Name()))
			if recursive {
				var recursivePath = filepath.Join(dir, file.Name())
				var dirpaths, err = e.ExploreDirectory(recursivePath, recursive)
				if err != nil {
					println("warning not found directory " + recursivePath)
					continue
				}

				paths = append(paths, dirpaths...)
			}
		}
	}

	return paths, nil
}
