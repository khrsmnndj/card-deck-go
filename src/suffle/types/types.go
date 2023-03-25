package types

import (
	"strconv"
)

type Card struct {
	Name string
}

type Cards []Card

func (c *Cards) Append(name string, length int) *Cards {	
	
	for i := 0; i < length; i++ {
		*c = append(*c,Card{Name: name + strconv.Itoa(i + 1)})
	}
	
	return c
}