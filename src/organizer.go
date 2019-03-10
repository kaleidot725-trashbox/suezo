package main

import (
	"io"
	"os"
	"path/filepath"
)

type Organizer struct {
	explorer Explorer
}

func (organizer Organizer) OriganizeByExtension(source string, destination string, copy bool) (err error) {
	var files, experr = organizer.explorer.ExploreFile(source, false)
	if experr != nil {
		return experr
	}

	var extmap = make(map[string]bool)
	var extuniq []string
	for _, item := range files {
		var ext = filepath.Ext(item)
		if !extmap[ext] {
			extmap[ext] = true
			extuniq = append(extuniq, ext[1:len(ext)])
		}
	}

	for _, item := range extuniq {
		var mkdirpath = filepath.Join(destination, item)
		os.Mkdir(mkdirpath, 0777)
	}

	for _, src := range files {
		var ext = filepath.Ext(src)
		var _, file = filepath.Split(src)
		var dst = filepath.Join(destination, ext[1:len(ext)], file)

		srcfile, err := os.Open(src)
		if err != nil {
			panic(err)
		}
		defer srcfile.Close()

		dstfile, err := os.Create(dst)
		if err != nil {
			panic(err)
		}
		defer dstfile.Close()

		_, err = io.Copy(dstfile, srcfile)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (organizer Organizer) getExtention(str string) (extention string) {
	var ext = filepath.Ext(str)
	return ext
}
