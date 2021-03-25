package mmap

import (
	"errors"
	"os"
)

// MaxFileSize max size of file
const MaxFileSize = 128 * (1 << 20)

// ErrOverRange map range over file size
// ErrNegativeSize negative size to map
// ErrZeroSize zero size to map
var (
	ErrOverRange    = errors.New("map range over file size")
	ErrNegativeSize = errors.New("map size can’t be negative")
	ErrZeroSize     = errors.New("map size can’t be zero")
)

// FileMemory is a structure to hold os.File and the memory slice
// use Data() to get the memory slice
type FileMemory struct {
	file *os.File
	m    []byte
}

// Data get the memory slice
func (fm *FileMemory) Data() []byte {
	return fm.m
}

// Map map the whole file into memory
// filePath : the location of the file
// this function open the file for memory mapping,
// it will close the file automatically
func Map(filePath string) (*FileMemory, error) {
	// open file
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return MapRegion(file, fileLength(file))
}

// MapRegion map the specified region of the file
func MapRegion(file *os.File, length int32) (fm *FileMemory, err error) {
	return mmap(file, length)
}

// Unmap unmap memory
func (fm *FileMemory) Unmap() (err error) {
	return fm.unmap()
}

// fileLength get the length of file
func fileLength(file *os.File) int32 {
	stat, _ := file.Stat()
	length := stat.Size()
	if length < 0 || length >= MaxFileSize {
		panic("file " + file.Name() + " is too big or corrupted")
	}
	if length == 0 {
		length = 1
	}
	return int32(length)
}

// checkRange check if the mapped area is out of range of the file
func checkRange(length int32, file *os.File) error {
	switch {
	case length > fileLength(file):
		return ErrOverRange
	case length < 0:
		return ErrNegativeSize
	case length == 0:
		return ErrZeroSize
	default:
		return nil
	}
}
