package main

// An interface can contain the name of one (or more) other interface(s), which
// is equivalent to explicitly enumerating the methods of the embedded interface
// in the containing interface. The effect is  as if we had written the embedded
// interface’s method signatures in the interface that embeds it.

// For example the standard library io package defines several interfaces, some
// of them are defined by embedding other interfaces.

// http://golang.org/src/io/io.go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader //embedded interface
	Closer //embedded interface
}

type ReadWriteCloser interface {
	Reader //embedded interface
	Writer //embedded interface
	Closer //embedded interface
}

func main() {

}
