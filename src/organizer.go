package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Organizer is
type Organizer struct {
	explorer Explorer
}

// OriganizeByExtension return err
func (o Organizer) OriganizeByExtension(source string, destination string, deleteOption bool) (err error) {
	var files, experr = o.explorer.ExploreFile(source, false)
	if experr != nil {
		return experr
	}

	var exts = o.createNoDuplicateExtArray(files)
	for _, item := range exts {
		var mkdirpath = filepath.Join(destination, item)
		os.Mkdir(mkdirpath, 0777)
	}

	for _, item := range files {
		organized := o.createOrganizedPath(item, destination)
		err = o.explorer.CopyFile(item, organized)
		if err != nil {
			fmt.Printf("copy %s %s\n", organized, err)
			break
		}

		err = o.explorer.Delete(item)
		if err != nil {
			fmt.Printf("delete %s %s\n", item, err)
			break
		}

		fmt.Printf("replace %s %s\n", item, organized)
	}

	// FIXME 適切なエラーを返す
	return err
}

func (o Organizer) createNoDuplicateExtArray(files []string) (exts []string) {
	var m = make(map[string]bool)
	for _, file := range files {
		ext := filepath.Ext(file)
		dirname := "none"
		if 0 < len(ext) {
			dirname = ext[1:len(ext)]
		}

		if !m[ext] {
			m[ext] = true

			exts = append(exts, dirname)
		}
	}

	return exts
}

func (o Organizer) createOrganizedPath(src string, dst string) (oraganized string) {
	ext := filepath.Ext(src)
	dirname := "none"
	if 0 < len(ext) {
		dirname = ext[1:len(ext)]
	}

	_, file := filepath.Split(src)
	oraganized = filepath.Join(dst, dirname, file)
	return oraganized
}
