package currency

type GrabsCurrency interface {
	GrabCbr() GrabCBR
}

type GrabCBR interface {
	GrabUsdEur() error
}
