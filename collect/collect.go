package collect

import (
	"errors"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Collecter base interface
type Collecter interface {
	Collect(length int) (string, error)
}

// Characters ...
type Characters struct {
	ContinueChar []int
	AlphabetChar []int
	SpecialChar  []string
	AllChar      []string
}

func (c *Characters) Collect(length int) (string, error) {
	for _, v := range c.AlphabetChar {
		n := sort.SearchInts(c.ContinueChar, v)
		if n < len(c.ContinueChar) && c.ContinueChar[n] == v {
			continue
		}
		c.AllChar = append(c.AllChar, string(byte(v)))
	}
	c.AllChar = append(c.AllChar, c.SpecialChar...)

	switch {
	case length <= 0:
		err := errors.New("length must be more than 0 but not more than 73")
		return "", err
	case length > 73:
		err := errors.New("length must be less than 73")
		return "", err
	default:
		rand.Seed(time.Now().UnixNano())
		for i := range c.AllChar {
			j := rand.Intn(i + 1)
			c.AllChar[i], c.AllChar[j] = c.AllChar[j], c.AllChar[i]
		}
	}
	result := strings.Join(c.AllChar[:length], "")

	return result, nil
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
			"!", "@", "#", "$", "%", "^", "&", "[", "]",
			"*", "_", "1", "2", "3", "4", "5", "6", "7",
			"8", "9", "10",
		},
	}
}
