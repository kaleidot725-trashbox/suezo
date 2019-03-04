package main

import "testing"

func TestExploreFile(t *testing.T) {
	var files []string = []string{
		"test/test1",
		"test/test2",
		"test/test3",
		"test/test4",
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
		"test/test1",
		"test/test2",
		"test/test3",
		"test/test4",
		"test/dir1/test5",
		"test/dir1/test6",
		"test/dir1/test7",
		"test/dir1/test8",
		"test/dir1/dir2/test9",
		"test/dir1/dir2/test10",
		"test/dir1/dir2/test11",
		"test/dir1/dir2/test12",
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

func contains(str string, array []string) bool {
	for _, item := range array {
		if str == item {
			return true
		}
	}
	return false
}
