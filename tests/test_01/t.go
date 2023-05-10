package main

import (
	"fmt"

	"github.com/AnimusPEXUS/gouuidtools"
)

func main() {
	{
		x := gouuidtools.NewUUIDNil()
		fmt.Println("nil:", x.Format(), x.GetVersion())
	}

	{
		x, err := gouuidtools.NewUUIDFromRandom()
		if err != nil {
			fmt.Println("random err:", err)
		} else {
			fmt.Println("random:", x.Format(), x.GetVersion())
		}
	}

}
