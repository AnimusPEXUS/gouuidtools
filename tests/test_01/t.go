package main

import (
	"fmt"

	"github.com/AnimusPEXUS/gouuidtools"
)

func main() {

	x := gouuidtools.NewUUIDNil()
	fmt.Println(x.Format())

}
