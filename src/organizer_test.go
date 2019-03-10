package main

import "testing"

func TestOrganizeByExtension(t *testing.T) {
	var explorer = Explorer{}
	var organizer = Organizer{explorer}
	var source = "../workspace"
	var destination = "../organized"
	var expectedPaths = []string{
		"organized/c/test1.c",
		"organized/cpp/test2.cpp",
		"organized/doc/test3.doc",
		"organized/jpg/test4.jpg",
	}

	err := organizer.OriganizeByExtension(source, destination)
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

func TestOrganizeByExtensionNotFoundSource(t *testing.T) {
	var organizer = Organizer{Explorer{}}
	var noSource = "./noSource"
	var destination = "./organized"

	err := organizer.OriganizeByExtension(noSource, destination)
	if err == nil {
		t.Error(err)
		return
	}

	println(err)
}

func TestOriganizeByExtensionNotFoundDestination(t *testing.T) {
	var organizer = Organizer{Explorer{}}
	var source = "./workspace"
	var noDestination = "./noDestination"

	err := organizer.OriganizeByExtension(source, noDestination)
	if err == nil {
		t.Error(err)
		return
	}

	println(err)
}
