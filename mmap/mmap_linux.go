package mmap

import (
	"os"
	"syscall"
)

// don't forget to close the file manually
func mmap(file *os.File, length int32) (fm *FileMemory, err error) {
	if err = checkRange(length, file); err != nil {
		return nil, err
	}
	m, err := syscall.Mmap(int(file.Fd()), 0, int(length),
		syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return
	}

	fm = &FileMemory{file, m}
	return
}

// the parameter size can't exceed the size of file
func (fm *FileMemory) Grow(size int32) error {
	if err := checkRange(size, fm.file); err != nil {
		return err
	}
	return fm.file.Truncate(int64(size))
}

// don't forget to close the file manually
func (fm *FileMemory) unmap() (err error) {
	if fm == nil {
		return nil
	}
	err = syscall.Munmap(fm.m)
	fm.m = nil
	return
}
