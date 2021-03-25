package mmapio

import (
	"errors"
	"io"
	"mmap-io-go/mmap"
)

// Reader save the pointer to read the file
type Reader struct {
	fileMemory *mmap.FileMemory
	cur        int
	lineNumber int
}

var (
	ErrNilReader = errors.New("can not invoking function of a nil reader")
)

// Read read the file data into the slice
// Returns the length of the data
func (reader *Reader) Read(p []byte) (n int, err error) {
	if reader == nil {
		return 0, ErrNilReader
	}
	data := reader.fileMemory.Data()
	r := data[reader.cur:]

	copy(p, r)
	if len(r) > len(p) {
		n = len(p)
	} else {
		n = len(r)
	}

	reader.cur += n
	return
}

// NewReader returns a reader that uses mmap io to read file
func NewReader(filePath string) (reader *Reader, err error) {
	if fm, err := mmap.Map(filePath); err != nil {
		return nil, err
	} else {
		return &Reader{fm, 0, 0}, nil
	}
}

// ReadAll read whole file and format to string
func (reader *Reader) ReadAll() (b []byte, err error) {
	if reader == nil {
		return nil, ErrNilReader
	}
	return reader.fileMemory.Data()[:], nil
}

// ReadLine returns one row of the file data
func (reader *Reader) ReadLine() (b []byte, err error) {
	if reader == nil {
		return nil, ErrNilReader
	}
	i := reader.cur
	data := reader.fileMemory.Data()

	// reach the end of file
	if len(data) == 0 || i >= len(data) {
		err = io.EOF
	}

	// detect the end of line
	for i < len(data)-1 && data[i] != '\n' {
		i++
	}

	// return a row of data
	b = data[reader.cur:i]

	// move current pointer to next line
	if i < len(data) {
		reader.cur = i + 1
	}
	reader.lineNumber++
	return
}

// LineNumber get current line number
func (reader *Reader) LineNumber() int {
	if reader == nil {
		panic(ErrNilReader)
	}
	return reader.lineNumber
}

// Close unmap file and release resources
func (reader *Reader) Close() error {
	if reader == nil {
		return nil
	}
	if err := reader.fileMemory.Unmap(); err != nil {
		return err
	}
	reader.fileMemory = nil
	return nil
}
