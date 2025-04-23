package main

import "fmt"

type Formatter interface {
	Format() (StringFormatada string)
}

func Format(f Formatter) (StringFormatada string) {

	return f.Format()

}

type PlainText struct {
	message string
}
type Bold struct {
	message string
}
type Code struct {
	message string
}

func (p PlainText) Format() (StringFormatada string) {
	return p.message
}

func (b Bold) Format() (StringFormatada string) {
	return "**" + b.message + "**"
}

func (c Code) Format() (StringFormatada string) {
	return "`" + c.message + "`"
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}

func main() {
	p1 := PlainText{message: "teste1"}
	b1 := Bold{message: "teste2"}
	c1 := Code{message: "teste3"}
	fmt.Println(Format(p1))
	fmt.Println(Format(b1))
	fmt.Println(Format(c1))
}
