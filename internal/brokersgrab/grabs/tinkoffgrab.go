package grabs

import (
	"context"
	"errors"
	"time"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/nile546/diplom/internal/models"
)

type TinkoffGrab struct {
	token string
}

func (t *TinkoffGrab) GetTinkoffStockDeals(token string) (*[]models.TinkoffOperation, error) {
	t.token = token
	return t.getTinkoffOperations()
}

func (t *TinkoffGrab) getTinkoffOperations() (*[]models.TinkoffOperation, error) {

	client := sdk.NewRestClient(t.token)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	operations, err := client.Operations(ctx, sdk.DefaultAccount, time.Now().AddDate(0, 0, -1000), time.Now(), "")
	if err != nil {
		return nil, err
	}

	if len(operations) <= 0 {
		return nil, errors.New("Нет сделок")
	}

	tinkoffOperations := &[]models.TinkoffOperation{}

	for i := len(operations) - 1; i >= 0; i-- {
		if operations[i].Status == "Done" {
			if operations[i].InstrumentType == "Stock" {
				if operations[i].OperationType == "Buy" || operations[i].OperationType == "Sell" {
					ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()
					instrument, err := client.InstrumentByFIGI(ctx, operations[i].FIGI)
					if err != nil {
						//TODO: ADD TO LOGGER
					}
					tinkoffOperation := &models.TinkoffOperation{
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

func currencyConvert(currency sdk.Currency) models.Currency {

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

func operationConvert(operation sdk.OperationType) models.Type {

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
