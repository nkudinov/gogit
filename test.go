package main

import (
	"fmt"
)

func main() {
	path := "/Users/nkudinov/lap-observer/.git/objects/ff/050e5f8dbf9998c6baee679d7c0ce0a0e49450"
	var content = gitgo.objectFromFile(path)
	fmt.Println(content)
}
