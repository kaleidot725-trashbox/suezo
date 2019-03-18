package main

import (
	"os"
	"testing"
)

func TestOrganizeByExtension(t *testing.T) {
	var explorer = Explorer{}
	var organizer = Organizer{explorer}
	var source = "./workspace"
	var destination = "./organized"
	var expectedPaths = []string{
		"organized/c/test1.c",
		"organized/cpp/test2.cpp",
		"organized/doc/test3.doc",
		"organized/jpg/test4.jpg",
	}

	os.Mkdir(destination, 0777)
	{
		err := organizer.OriganizeByExtension(source, destination, false)
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
	explorer.Delete(destination)
}

func TestOrganizeByExtention_削除オプション(t *testing.T) {
	explorer := Explorer{}
	organizer := Organizer{explorer}
	source := "./workspace"
	delete := "./delete"
	destination := "./organized"

	os.Mkdir(destination, 0777)
	{
		err := explorer.CopyDirectory(source, delete)
		if err != nil {
			t.Error(err)
			return
		}

		err = organizer.OriganizeByExtension(delete, destination, true)
		if err != nil {
			t.Error(err)
		}

		bret := explorer.Exists(delete)
		if bret {
			t.Error(err)
			return
		}
	}
}

func TestOrganizeByExtensionNotFoundSource(t *testing.T) {
	var organizer = Organizer{Explorer{}}
	var noSource = "./noSource"
	var destination = "./organized"

	err := organizer.OriganizeByExtension(noSource, destination, false)
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

	err := organizer.OriganizeByExtension(source, noDestination, false)
	if err == nil {
		t.Error(err)
		return
	}

	println(err)
}
