package printer

import (
	"fmt"
	"io"
	"os"
)

type Printer interface {
	SetOutput(w io.Writer)

	Print(a ...interface{}) (n int, err error)
	Printf(format string, a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
}

type SimplePrinter struct {
	w io.Writer
}

func NewSimplePrinter(w io.Writer) *SimplePrinter {
	return &SimplePrinter{
		w: w,
	}
}

func (p *SimplePrinter) SetOutput(w io.Writer) {
	p.w = w
}

func (p SimplePrinter) Print(a ...interface{}) (n int, err error) {
	if p.w == nil {
		p.w = os.Stdout
	}
	return fmt.Fprint(p.w, a...)
}

func (p SimplePrinter) Printf(format string, a ...interface{}) (n int, err error) {
	if p.w == nil {
		p.w = os.Stdout
	}
	return fmt.Fprintf(p.w, format, a...)
}

func (p SimplePrinter) Println(a ...interface{}) (n int, err error) {
	if p.w == nil {
		p.w = os.Stdout
	}
	return fmt.Fprintln(p.w, a...)
}
