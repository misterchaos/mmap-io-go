package mmap

import (
	"os"
	"testing"
)

const testFile = "../testdata"

func TestMap(t *testing.T) {
	if fm, err := Map(testFile); err != nil {
		t.Errorf(err.Error())
	} else {
		defer fm.Unmap()
		if testFile != fm.file.Name() {
			t.Errorf("testFile is not the same")
		}
	}
}

func TestMapRegion(t *testing.T) {
	// open file
	file, err := os.OpenFile(testFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	defer file.Close()
	// mmap
	length := fileLength(file)
	fm, err := MapRegion(file, length)
	defer fm.Unmap()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(fm.m) != int(length) {
		t.Errorf("file memory's length not match")
	}

	// test length over file size
	length = length + 1
	fm2, err := MapRegion(file, length)
	defer fm2.Unmap()
	if err == nil {
		t.Errorf("length over file size")
	}

	// test length is negative
	length = -1
	fm3, err := MapRegion(file, length)
	defer fm3.Unmap()
	if err == nil {
		t.Errorf("length is negative")
	}

	// test zero length
	length = 0
	fm4, err := MapRegion(file, int32(length))
	defer fm4.Unmap()
	if err == nil {
		t.Errorf("length is zero")
	}
}
