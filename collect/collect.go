package collect

import "fmt"

// Collecter base interface
type Collecter interface {
	Collect(length int) string
}

// Characters ...
type Characters struct {
	ContinueChar []uint
	AlphabetChar []uint
	SpecialChar  []string
}

func (c *Characters) Collect(length int) string {
	return fmt.Sprint("Method interface Collect ", length)
}

func getRange(start uint, end uint) (result []uint) {
	result = make([]uint, 0, end-start)
	for value := start; value <= end; value++ {
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
