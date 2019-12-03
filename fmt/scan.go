package fmt

import (
	"io"
)

type stdin interface {
	// from os.Stdin
	Scan(a ...interface{}) (n int, err error)
	Scanln(a ...interface{}) (n int, err error)
	Scanf(format string, a ...interface{}) (n int, err error)

	// from string
	Sscan(str string, a ...interface{}) (n int, err error)
	Sscanln(str string, a ...interface{}) (n int, err error)
	Sscanf(str string, format string, a ...interface{}) (n int, err error)


	// from io.Reader
	Fscan(r io.Reader, a ...interface{}) (n int, err error)
	Fscanln(r io.Reader, a ...interface{}) (n int, err error)
	Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
}

type stdout interface {
	// to io.Writer
	Fprint(w io.Writer, a ...interface{}) (n int, err error)
	Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

	// to io.Stdout
	Print(a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
	Printf(format string, a ...interface{}) (n int, err error)

	// to string
	Sprint(a ...interface{}) string
	Sprintln(a ...interface{}) string
	Sprintf(format string, a ...interface{}) string
}
