package main

import (
	"fmt"
	"github.com/Beloslav13/char-collect/collect"
)

func getCollect(c collect.Collecter) (string, error) {
	//switch v := c.(type) {
	switch c.(type) {
	case *collect.Characters:
		characters := c.(*collect.Characters)
		chars, err := characters.Collect(12)
		if err != nil {
			return "", err
		}
		return chars, err
	default:
		fmt.Println("Default switch...")
		return "", nil
	}
}

func main() {
	var c collect.Collecter
	c = collect.NewCharacters()
	chars, err := getCollect(c)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Println(chars)
	}
}
