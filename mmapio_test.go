package mmapio

import (
	"strconv"
	"testing"
)

var testFile = "./testdata"

func TestNewReader(t *testing.T) {
	reader, err := NewReader(testFile)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_ = reader.Close()
}

func TestReader_ReadLine(t *testing.T) {
	if reader, err := NewReader(testFile); err == nil {
		if str, err := reader.ReadLine(); err == nil {
			t.Log("read the first line : " + string(str))
		} else {
			t.Error(err)
		}
		_ = reader.Close()
	} else {
		t.Error(err)
	}
}

func TestReader_ReadAll(t *testing.T) {
	if reader, err := NewReader(testFile); err == nil {
		if str, err := reader.ReadAll(); err == nil {
			t.Log("read whole file :\n" + string(str))
		} else {
			t.Error(err)
		}
		_ = reader.Close()
	} else {
		t.Error(err)
	}
}

func TestReader_LineNumber(t *testing.T) {
	if reader, err := NewReader(testFile); err == nil {
		_, _ = reader.ReadLine()
		if lineNumber := reader.LineNumber(); lineNumber != 1 {
			t.Error("line number incorrect ,should be 1 ,but is " + strconv.Itoa(lineNumber))
		}
		_ = reader.Close()
	} else {
		t.Error(err)
	}
}
