package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
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

func (e Explorer) CopyFile(source string, destination string) (err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return err
}

func (e Explorer) CopyDirectory(source string, destination string) (err error) {
	err = copy.Copy(source, destination)
	return err
}

func (e Explorer) Delete(target string) (err error) {
	err = os.RemoveAll(target)
	return err
}

func (e Explorer) Exists(path string) (b bool) {
	file, err := os.Open(path)
	defer file.Close()
	return err == nil
}
