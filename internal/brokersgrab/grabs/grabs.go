package grabs

import (
	"github.com/nile546/diplom/internal/brokersgrab"
)

type Grabs struct {
	tinkoffGrab *TinkoffGrab
}

func New() *Grabs {
	return &Grabs{}
}

func (g *Grabs) TinkoffGrab() brokersgrab.TinkoffGrab {
	if g.tinkoffGrab != nil {
		return g.tinkoffGrab
	}

	g.tinkoffGrab = &TinkoffGrab{}

	return g.tinkoffGrab
}
