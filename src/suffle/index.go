package suffle

import (
	"github.com/khrsmnndj/card-deck-go/src/suffle/types"
)

func GetCard(names string) types.Card{
	card := types.Card{names}
	return card
}