package main

import "testing"

func TestExploreFile(t *testing.T) {
	var explorer = Explorer{}
	var directory = "./workspace"
	var noRecursive = false
	var files []string = []string{
		"workspace/test1.c",
		"workspace/test2.cpp",
		"workspace/test3.doc",
		"workspace/test4.jpg",
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
	var directory = "./workspace"
	var recursive = true
	var recursiveFiles = []string{
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
	var directory = "./workspace"
	var noRecursive = false
	var expectedDirectories = []string{
		"workspace/dir1",
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
	var directory = "./workspace"
	var recurisive = true
	var recursiveDirectories = []string{
		"workspace/dir1",
		"workspace/dir1/dir2",
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
