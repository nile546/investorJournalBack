package grabs

import (
	"context"
	"time"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/nile546/diplom/internal/models"
)

type TinkoffGrab struct {
	token    string
	grabDate time.Time
}

func (t *TinkoffGrab) GetTinkoffStockDeals(token string, grabDate time.Time) (*[]models.BrokerOperation, error) {
	t.token = token
	t.grabDate = grabDate
	return t.getTinkoffOperations()
}

func (t *TinkoffGrab) getTinkoffOperations() (*[]models.BrokerOperation, error) {

	client := sdk.NewRestClient(t.token)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	operations, err := client.Operations(ctx, sdk.DefaultAccount, t.grabDate, time.Now(), "")
	if err != nil {
		return nil, err
	}

	if len(operations) <= 0 {
		return nil, nil
	}

	tinkoffOperations := &[]models.BrokerOperation{}

	for i := len(operations) - 1; i >= 0; i-- {
		if operations[i].Status == "Done" {
			if operations[i].InstrumentType == "Stock" {
				if operations[i].OperationType == "Buy" || operations[i].OperationType == "Sell" {
					ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()
					instrument, err := client.InstrumentByFIGI(ctx, operations[i].FIGI)
					if err != nil {
						continue
					}
					tinkoffOperation := &models.BrokerOperation{
						ISIN:      instrument.ISIN,
						Currency:  currencyConvert(operations[i].Currency),
						Quantity:  operations[i].QuantityExecuted,
						DateTime:  operations[i].DateTime,
						Price:     int64(operations[i].Price * 100),
						Operation: operationConvert(operations[i].OperationType),
					}
					*tinkoffOperations = append(*tinkoffOperations, *tinkoffOperation)
				}
			}
		}
	}

	return tinkoffOperations, nil

}

func currencyConvert(currency sdk.Currency) models.Currencies {

	switch currency {
	case "USD":
		{
			return 1
		}
	case "EUR":
		{
			return 2
		}
	case "RUB":
		{
			return 3
		}
	}
	return 0
}

func operationConvert(operation sdk.OperationType) models.DealTypes {

	switch operation {
	case "Buy":
		{
			return 1
		}
	case "Sell":
		{
			return 2
		}

	}
	return 0
}
