package currency

type GrabsCurrency interface {
	GrabCBR() GrabCBR
}

type GrabCBR interface {
	GrabUsdEur()
}
