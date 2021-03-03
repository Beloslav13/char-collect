package collect

import "fmt"

type BaseCharacters interface {
	Collect(length int) string
}

type Characters struct {
	ContinueChar []uint
	AlphabetChar []uint
	SpecialChar  []string
}

func getRange(start uint, end uint) (result []uint) {
	result = make([]uint, 0, end-start)
	for value := start; value <= end; value++ {
		result = append(result, value)
	}
	return result
}

func (c *Characters) Collect(length int) string {
	return fmt.Sprint("Method interface Collect ", length)
}

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
