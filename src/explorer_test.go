package main

import (
	"os"
	"testing"
)

func TestExploreFile(t *testing.T) {
	explorer := Explorer{}
	directory := "./workspace"
	noRecursive := false
	files := []string{
		"workspace/test1.c",
		"workspace/test2.cpp",
		"workspace/test3.doc",
		"workspace/test4.jpg",
	}

	paths, err := explorer.ExploreFile(directory, noRecursive)
	if err != nil {
		t.Error(err)
	}

	for _, item := range files {
		bret := contains(item, paths)
		if !bret {
			t.Error("not found " + item)
		}
	}

	if len(paths) != len(files) {
		t.Error("count is not equal")
	}
}

func TestExploreFileRecursive(t *testing.T) {
	explorer := Explorer{}
	directory := "./workspace"
	recursive := true
	recursiveFiles := []string{
		"workspace/test1.c",
		"workspace/test2.cpp",
		"workspace/test3.doc",
		"workspace/test4.jpg",
		"workspace/dir1/test5.wav",
		"workspace/dir1/test6.txt",
		"workspace/dir1/test7.mp3",
		"workspace/dir1/test8.zip",
		"workspace/dir1/dir2/test9.xls",
		"workspace/dir1/dir2/test10.png",
		"workspace/dir1/dir2/test11.log",
		"workspace/dir1/dir2/test12.dat",
	}

	paths, err := explorer.ExploreFile(directory, recursive)
	if err != nil {
		t.Error(err)
	}

	for _, item := range recursiveFiles {
		bret := contains(item, paths)
		if !bret {
			t.Error("not found " + item)
		}
	}

	if len(paths) != len(recursiveFiles) {
		t.Error("count is not equal")
	}
}

func TestExploreFileNotFoundDirectory(t *testing.T) {
	explorer := Explorer{}
	noDirectory := "./noDir"
	paths, err := explorer.ExploreFile(noDirectory, true)
	if paths != nil && err == nil {
		t.Error(err)
	}
}

func TestExploreDirectory(t *testing.T) {
	explorer := Explorer{}
	directory := "./workspace"
	noRecursive := false
	expectedDirectories := []string{
		"workspace/dir1",
	}

	paths, err := explorer.ExploreDirectory(directory, noRecursive)
	if err != nil {
		t.Error(err)
	}

	for _, item := range expectedDirectories {
		bret := contains(item, paths)
		if !bret {
			t.Error("not found " + item)
		}
	}

	if len(paths) != len(expectedDirectories) {
		t.Error("count is not equal")
	}
}

func TestExploreDirectoryRecursive(t *testing.T) {
	explorer := Explorer{}
	directory := "./workspace"
	recurisive := true
	recursiveDirectories := []string{
		"workspace/dir1",
		"workspace/dir1/dir2",
	}

	paths, err := explorer.ExploreDirectory(directory, recurisive)
	if err != nil {
		t.Error(err)
	}

	for _, item := range recursiveDirectories {
		bret := contains(item, paths)
		if !bret {
			t.Error("not found " + item)
		}
	}

	if len(paths) != len(recursiveDirectories) {
		t.Error("count is not equal")
	}
}

func TestExploreDirecotryNotFoundDirectory(t *testing.T) {
	var explorer = Explorer{}
	var noDirectory = "./nodir"
	var paths, err = explorer.ExploreDirectory(noDirectory, true)
	if paths != nil && err == nil {
		t.Error(err)
	}
}

func TestExploreExists(t *testing.T) {
	explorer := Explorer{}
	directory := "./workspace"
	file := "./workspace/test1.c"

	b := explorer.Exists(directory)
	if !b {
		t.Error("not found error")
	}

	b = explorer.Exists(file)
	if !b {
		t.Error("not found error")
	}
}

func TestExplorerExists_ない場合(t *testing.T) {
	explorer := Explorer{}
	noDirectory := "./nodir"
	noFile := "./nodir/nofile.txt"

	b := explorer.Exists(noDirectory)
	if b {
		t.Error("found error")
	}

	b = explorer.Exists(noFile)
	if b {
		t.Error("found error")
	}
}

func TestCopyFileとDelete(t *testing.T) {
	explorer := Explorer{}
	source := "./workspace/test1.c"
	delete := "./delete/test1.c"

	os.Mkdir("./delete", 0777)
	{
		err := explorer.CopyFile(source, delete)
		if err != nil {
			t.Error(err)
			return
		}

		copyed := explorer.Exists(delete)
		if !copyed {
			t.Error(err)
			return
		}

		err = explorer.Delete(delete)
		if err != nil {
			t.Error(err)
			return
		}

		copyed = explorer.Exists(delete)
		if copyed {
			t.Error(err)
			return
		}
	}
	explorer.Delete("./delete")
}

func TestCopyDirectoryとDelete(t *testing.T) {
	explorer := Explorer{}
	source := "./workspace/"
	delete := "./delete"
	deleteFiles := []string{
		"delete/test1.c",
		"delete/test2.cpp",
		"delete/test3.doc",
		"delete/test4.jpg",
		"delete/dir1/test5.wav",
		"delete/dir1/test6.txt",
		"delete/dir1/test7.mp3",
		"delete/dir1/test8.zip",
		"delete/dir1/dir2/test9.xls",
		"delete/dir1/dir2/test10.png",
		"delete/dir1/dir2/test11.log",
		"delete/dir1/dir2/test12.dat",
	}

	err := explorer.CopyDirectory(source, delete)
	if err != nil {
		t.Error(err)
		return
	}

	exists := explorer.Exists(delete)
	if !exists {
		t.Error("doesn't create directory")
		return
	}

	existsFiles, err := explorer.ExploreFile(delete, true)
	if err != nil {
		t.Error(err)
		return
	}

	for _, file := range existsFiles {
		bret := contains(file, deleteFiles)
		if !bret {
			t.Error("doens't copy " + file)
		}
	}

	err = explorer.Delete(delete)
	if err != nil {
		t.Error(err)
		return
	}

	exists = explorer.Exists(delete)
	if exists {
		t.Error("doesn't delete directory")
		return
	}
}
