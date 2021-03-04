package main

import (
	"char-collect/collect"
	"fmt"
)

func getCollect(c collect.Collecter) {
	//switch v := c.(type) {
	switch c.(type) {
	case *collect.Characters:
		characters := c.(*collect.Characters)
		fmt.Println(characters.Collect(12))

	}
}

func main() {
	var c collect.Collecter
	c = collect.NewCharacters()
	getCollect(c)
}
