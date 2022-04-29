package main

import (
	"errors"
	"fmt"

	"github.com/SAMCHEE/practice_git/util"
)

func main() {
	fmt.Println("Hello git")

	util.PanicIfErr(nil)
	util.PanicIfErr(errors.New("Hello error"))

	fmt.Println("Bye")
}
