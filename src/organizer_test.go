package main

import "testing"

func TestOrganizeByExtension(t *testing.T) {
	var explorer = Explorer{}
	var organizer = Organizer{Explorer{}}
	var source = "./test"
	var destination = "./organized"
	var expectedPaths = []string{
		"organized/c/test1.c",
		"organized/cpp/test2.cpp",
		"organized/doc/test3.doc",
		"organized/jpg/test4.jpg",
	}

	err := organizer.OriganizeByExtension(source, destination, true)
	if err != nil {
		t.Error(err)
		return
	}

	actualPaths, err := explorer.ExploreFile(destination, true)
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range expectedPaths {
		if !contains(item, actualPaths) {
			t.Error("not found " + item)
			return
		}
	}
}
