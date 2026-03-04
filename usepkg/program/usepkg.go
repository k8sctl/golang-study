package main

import (
	"golang/usepkg/custompkg"
	"golang/usepkg/exinit"
)

func main() {
	custompkg.PrintCustom()
	exinit.PrintD()
}

/*
[출력 결과]
f() d: 4
f() d: 5
exinit.init function 6
This is custom package
d: 6
d: 6
*/
