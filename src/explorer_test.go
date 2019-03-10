package main

import "testing"

func TestExploreFile(t *testing.T) {
	var explorer = Explorer{}
	var directory = "./test"
	var noRecursive = false
	var files []string = []string{
		"test/test1.c",
		"test/test2.cpp",
		"test/test3.doc",
		"test/test4.jpg",
	}

	var paths, err = explorer.ExploreFile(directory, noRecursive)
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
	var explorer = Explorer{}
	var directory = "./test"
	var recursive = true
	var recursiveFiles = []string{
		"test/test1.c",
		"test/test2.cpp",
		"test/test3.doc",
		"test/test4.jpg",
		"test/dir1/test5.wav",
		"test/dir1/test6.txt",
		"test/dir1/test7.mp3",
		"test/dir1/test8.zip",
		"test/dir1/dir2/test9.xls",
		"test/dir1/dir2/test10.png",
		"test/dir1/dir2/test11.log",
		"test/dir1/dir2/test12.dat",
	}

	var paths, err = explorer.ExploreFile(directory, recursive)
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
	var explorer = Explorer{}
	var noDirectory = "./noDir"
	var paths, err = explorer.ExploreFile(noDirectory, true)
	if paths != nil && err == nil {
		t.Error(err)
	}
}

func TestExploreDirectory(t *testing.T) {
	var explorer = Explorer{}
	var directory = "./test"
	var noRecursive = false
	var expectedDirectories = []string{
		"test/dir1",
	}

	var paths, err = explorer.ExploreDirectory(directory, noRecursive)
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
	var explorer = Explorer{}
	var directory = "./test"
	var recurisive = true
	var recursiveDirectories = []string{
		"test/dir1",
		"test/dir1/dir2",
	}

	var paths, err = explorer.ExploreDirectory(directory, recurisive)
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
