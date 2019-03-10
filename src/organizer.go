package main

import (
	"io"
	"os"
	"path/filepath"
)

// Organizer is
type Organizer struct {
	explorer Explorer
}

// OriganizeByExtension return err
func (organizer Organizer) OriganizeByExtension(source string, destination string) (err error) {
	var files, experr = organizer.explorer.ExploreFile(source, false)
	if experr != nil {
		return experr
	}

	var exts = organizer.createNoDuplicateExtArray(files)
	for _, item := range exts {
		var mkdirpath = filepath.Join(destination, item)
		os.Mkdir(mkdirpath, 0777)
	}

	for _, item := range files {
		oraganized := organizer.createOrganizedPath(item, destination)
		err = organizer.copy(item, oraganized)
		if err != nil {
			break
		}
	}

	return err
}

func (organizer Organizer) copy(src string, dst string) (err error) {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
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

func (organizer Organizer) createNoDuplicateExtArray(files []string) (exts []string) {
	var m = make(map[string]bool)
	for _, file := range files {
		var ext = filepath.Ext(file)
		if !m[ext] {
			m[ext] = true
			exts = append(exts, ext[1:len(ext)])
		}
	}

	return exts
}

func (organizer Organizer) createOrganizedPath(src string, dst string) (oraganized string) {
	ext := filepath.Ext(src)
	_, file := filepath.Split(src)
	oraganized = filepath.Join(dst, ext[1:len(ext)], file)
	return oraganized
}
