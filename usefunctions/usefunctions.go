package main

import (
    "fmt"
    "github.com/jinkyaw/gocheckpackage/greetings"
	"github.com/jinkyaw/gocheckpackage/reversestring"
)

func main() {
	fmt.Println(greetings.Hello("Chetna"))
    fmt.Println(reversestring.ReverseRunes("Chetna"))
}
