package grabscurrency

import currency "github.com/nile546/diplom/internal/apiserver/currencygrab"

type GrabsCurrency struct {
	grabCBR *GrabCBR
}

func New() *GrabsCurrency {
	return &GrabsCurrency{}
}

func (g *GrabsCurrency) GrabCBR() currency.GrabCBR {

	if g.grabCBR != nil {
		return g.grabCBR
	}

	g.grabCBR = &GrabCBR{}

	return g.grabCBR

}
