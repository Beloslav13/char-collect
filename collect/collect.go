package collect

import "fmt"

// Collecter base interface
type Collecter interface {
	Collect(length int) string
}

// Characters ...
type Characters struct {
	ContinueChar map[byte]byte
	AlphabetChar map[byte]byte
	SpecialChar  []string
	AllChar      []string
}

func (c *Characters) Collect(length int) string {
	for key, _ := range c.AlphabetChar {
		fmt.Println(string(c.AlphabetChar[key]))
		if _, exists := c.ContinueChar[key]; exists {
			continue
		}
		c.AllChar = append(c.AllChar, string(c.AlphabetChar[key]))
	}

	return fmt.Sprint("Method interface Collect ", length, c.AllChar, c.AllChar[51])
}

func getRange(start byte, end byte) (result map[byte]byte) {
	result = make(map[byte]byte, end-start)
	//result = make([]byte, 0, end-start)
	// todo: элемент последний end
	for value := start; value <= end; value++ {
		//result = append(result, value)
		result[value] = value
	}
	fmt.Println(result)
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
