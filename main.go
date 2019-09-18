package main

import (
	pkg1 "github.com/johnrichardrinehart/go_modules_test/mod1/pkg1"
	pkg2 "github.com/johnrichardrinehart/go_modules_test/mod2/pkg2"
	"fmt"
)

func main() {
	fmt.Println(pkg1.PackageAndVersion())
	fmt.Println(pkg2.PackageAndVersion())
}
