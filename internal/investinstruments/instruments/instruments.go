package instruments

import (
	"io/ioutil"
	"strings"

	"github.com/nile546/diplom/internal/investinstruments"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type Instruments struct {
	stockinstrumnet  *Stockinstrument
	bankinstruments  *Bankinstruments
	cryptoinstrumnet *Cryptoinstrument
}

var log *logrus.Logger

func New(l *logrus.Logger) *Instruments {
	log = l
	return &Instruments{}
}

func (i *Instruments) Stocks() investinstruments.Stockinstrument {

	if i.stockinstrumnet != nil {
		return i.stockinstrumnet
	}

	i.stockinstrumnet = &Stockinstrument{}

	return i.stockinstrumnet
}

func (i *Instruments) Cryptos() investinstruments.Cryptoinstrument {

	if i.cryptoinstrumnet != nil {
		return i.cryptoinstrumnet
	}

	i.cryptoinstrumnet = &Cryptoinstrument{}

	return i.cryptoinstrumnet
}

func (i *Instruments) Banks() investinstruments.Bankinstruments {

	if i.bankinstruments != nil {
		return i.bankinstruments
	}

	i.bankinstruments = &Bankinstruments{}

	return i.bankinstruments
}

func convertWin1251toUTF8(cs string) (string, error) {
	sr := strings.NewReader(cs)
	tr := transform.NewReader(sr, charmap.Windows1251.NewDecoder())
	buf, err := ioutil.ReadAll(tr)
	if err != err {
		return "", err
	}
	return string(buf), nil
}
