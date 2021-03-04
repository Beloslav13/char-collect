package main

import (
	"fmt"
	"github.com/Beloslav13/char-collect/collect"
)

func getCollect(c collect.Collecter) {
	//switch v := c.(type) {
	switch c.(type) {
	case *collect.Characters:
		characters := c.(*collect.Characters)
		fmt.Println(characters.Collect(12))
	default:
		fmt.Println("Default switch...")
	}
}

func main() {
	var c collect.Collecter
	c = collect.NewCharacters()
	getCollect(c)
}
