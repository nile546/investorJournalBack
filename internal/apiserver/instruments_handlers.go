package apiserver

type instrumentsConfig struct {
	spbExchangeUrl string
	mskStocksUrl   string
	bankiUrl       string
	cryptoUrl      string
	cryptoKey      string
}

/*func updateInstruments(hour, min, sec int, f func()) error {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	now := time.Now().Local()
	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), hour, min, sec, 0, loc)
	if firstCallTime.Before(now) {
		firstCallTime = firstCallTime.Add(time.Hour * 24)
	}

	duration := firstCallTime.Sub(time.Now().Local())

	go func() {
		time.Sleep(duration)
		for {
			f()
			time.Sleep(time.Hour * 24)
		}
	}()

	return nil
}*/

func updateInstruments(c *instrumentsConfig) {

}
