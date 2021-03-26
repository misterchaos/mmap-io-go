# Mmap I/O for Golang

A portable I/O implement base on mmap for golang.

## Overview

Use it like `bufio` in golang.

- `NewReader(filePath string)` returns a reader that uses mmap io to read file
- `Read(p []byte) (n int, err error)` read the file data into the slice
- `ReadAll()(b []byte, err error)` read whole file and format to string
- `ReadLine() (b []byte, err error)` returns one row of the file data
- `LineNumber() int` get current line number
- `Close() error`  unmap file and release resources

## Usage

### Basic Usage

Create a reader from a filepath

```go
reader, err := NewReader(testFile)
if err != nil{
	return
}
defer reader.Close()
```

Read data into slice

```go
b := make([]byte, 20)

if n, err := reader.Read(b); err == nil {
	fmt.Printf("read " + strconv.Itoa(n) + " bytes data from file" + " : " + string(b))
}
```

The output of the above code is as follows

```bash
Welcome to mmap io !
```

### Read the whole file at once

You can use `ReadAll()` to read it at once

```go
if reader, err := NewReader(testFile); err == nil {
	defer reader.Close()
	if b, err := reader.ReadAll(); err == nil {
		fmt.Printf(string(b))	
	}
}
```

The output of the above code is as follows

```shell
Welcome to mmap io ! by Yuchao Huang @misterchaos
 _          _ _                                         _
| |__   ___| | | ___    _ __ ___  _ __ ___   __ _ _ __ (_) ___
| '_ \ / _ \ | |/ _ \  | '_ ` _ \| '_ ` _ \ / _` | '_ \| |/ _ \
| | | |  __/ | | (_) | | | | | | | | | | | | (_| | |_) | | (_) |
|_| |_|\___|_|_|\___/  |_| |_| |_|_| |_| |_|\__,_| .__/|_|\___/
                                                 |_|
```

## ToDo

- Add file writing and locking support
- Implements `io.Reader` interface to be compatible with `bufio`
- Implements a mapping of the specified region

## License

The code is open source using GPL3 protocol. If you need to use the code, please follow the relevant provisions of CPL3
protocol.

## Authors

- Yuchao Huang [@misterchaos](https://github.com/misterchaos/) - Original Author

