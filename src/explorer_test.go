package main

import "testing"

func TestExploreFile(t *testing.T) {
	var files []string = []string{
		"test/test1.c",
		"test/test2.cpp",
		"test/test3.doc",
		"test/test4.jpg",
	}

	var explorer = Explorer{}
	var paths = explorer.ExploreFile("./test", false)

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

	var explorer = Explorer{}
	var paths = explorer.ExploreFile("./test", true)

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

func TestExploreDirectory(t *testing.T) {
	var directories = []string{
		"test/dir1",
	}

	var explorer = Explorer{}
	var paths = explorer.ExploreDirectory("./test", false)

	for _, item := range directories {
		bret := contains(item, paths)
		if !bret {
			t.Error("not found " + item)
		}
	}

	if len(paths) != len(directories) {
		t.Error("count is not equal")
	}
}

func TestExploreDirectoryRecursive(t *testing.T) {
	var recursiveDirectories = []string{
		"test/dir1",
		"test/dir1/dir2",
	}

	var explorer = Explorer{}
	var paths = explorer.ExploreDirectory("./test", true)

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

func contains(str string, array []string) bool {
	for _, item := range array {
		if str == item {
			return true
		}
	}
	return false
}
