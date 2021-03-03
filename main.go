package main

import (
	"char-collect/collect"
	"fmt"
)

func main() {
	//c := collect.NewCharacters()
	var a collect.BaseCharacters
	a = collect.NewCharacters()
	fmt.Println(a.Collect(10))
	fmt.Println("====")
	//characters := collect.NewCharacters()
	//fmt.Println(characters.ContinueChar)
	//fmt.Println(characters.AlphabetChar)
	//fmt.Println(characters.SpecialChar)
	//
	//fmt.Println(characters.Collect(10))
	//
	//stra := "the spice must flow"
	//byts := []byte(stra)
	//strb := string(byts)
	//fmt.Println(byts)
	//fmt.Println(strb)

}
