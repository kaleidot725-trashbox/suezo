package main

import (
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
		oraganized := o.createOrganizedPath(item, destination)
		err = o.explorer.CopyFile(item, oraganized)
		if err != nil {
			return err
		}
	}

	if deleteOption {
		err = o.explorer.Delete(source)
		if err != nil {
			return err
		}
	}

	err = nil
	return err
}

func (o Organizer) createNoDuplicateExtArray(files []string) (exts []string) {
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

func (o Organizer) createOrganizedPath(src string, dst string) (oraganized string) {
	ext := filepath.Ext(src)
	_, file := filepath.Split(src)
	oraganized = filepath.Join(dst, ext[1:len(ext)], file)
	return oraganized
}
