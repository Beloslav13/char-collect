package collect

import (
	"fmt"
	"sort"
)

// Collecter base interface
type Collecter interface {
	Collect(length int) string
}

// Characters ...
type Characters struct {
	ContinueChar []int
	AlphabetChar []int
	SpecialChar  []string
	AllChar      []string
}

func (c *Characters) Collect(length int) string {
	for _, v := range c.AlphabetChar {
		n := sort.SearchInts(c.ContinueChar, v)
		if n < len(c.ContinueChar) && c.ContinueChar[n] == v {
			continue
		}
		c.AllChar = append(c.AllChar, string(byte(v)))
	}
	c.AllChar = append(c.AllChar, c.SpecialChar...)
	return fmt.Sprint("Method interface Collect ", length, c.AllChar)
}

func getRange(start int, end int) (result []int) {
	result = make([]int, 0, end-start)
	for value := start; value < end; value++ {
		result = append(result, value)
	}
	return result
}

// NewCharacters init object
func NewCharacters() *Characters {
	return &Characters{
		ContinueChar: getRange(91, 97),
		AlphabetChar: getRange(65, 123),
		SpecialChar: []string{
			"!", "@", "#", "$", "%", "^", "&",
			"*", "_", "1", "2", "3", "4", "5",
			"6", "7", "8", "9", "10",
		},
	}
}
