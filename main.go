package main

import (
	"fmt"

	"github.com/nkudinov/gogit/gogit"
)

func main() {
	path := "/Users/nkudinov/lap-observer/.git/objects/ff/050e5f8dbf9998c6baee679d7c0ce0a0e49450"
	content, err := gogit.ObjectFromFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// print content of file
	fmt.Println(content)

}
