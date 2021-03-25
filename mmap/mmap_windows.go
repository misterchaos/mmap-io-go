package mmap

import (
	"os"
	"syscall"
	"unsafe"
)

// don't forget to close the file manually
func mmap(file *os.File, length int32) (fm *FileMemory, err error) {
	if err = checkRange(length, file); err != nil {
		return nil, err
	}

	h, err := syscall.CreateFileMapping(syscall.Handle(file.Fd()), nil,
		syscall.PAGE_READWRITE, 0, uint32(length), nil)
	if h == 0 {
		return
	}

	addr, err := syscall.MapViewOfFile(h, syscall.FILE_MAP_READ,
		0, 0, uintptr(length))
	if addr == 0 {
		return
	}

	if err = syscall.CloseHandle(syscall.Handle(h)); err != nil {
		return
	}

	// Convert to a byte array.
	var sl = struct {
		addr uintptr
		len  int
		cap  int
	}{addr, int(length), int(length)}

	// Use unsafe to turn sl into a []byte.
	fm = &FileMemory{file, *(*[]byte)(unsafe.Pointer(&sl))}
	return
}

// don't forget to close the file manually
func (fm *FileMemory) unmap() (err error) {
	if fm == nil {
		return nil
	}
	addr := (uintptr)(unsafe.Pointer(&fm.m[0]))
	return syscall.UnmapViewOfFile(addr)
}
