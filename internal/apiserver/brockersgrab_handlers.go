package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/nile546/diplom/internal/models"
)

func (s *server) GetTinkoffStockDeals(w http.ResponseWriter, r *http.Request) {

	type request struct {
		token string
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	tinkoff_operations, err := s.brokersGrab.TinkoffGrab().GetTinkoffStockDeals(req.token)
	if err != nil {
		s.respond(w, err.Error())
		return
	}

	stock_deals := &[]models.StockDeal{}

	for i, operation := range *tinkoff_operations {
		if operation.Operation == "Buy" {

			stockdealID, err := s.repository.StockDeal().GetStockDealsIDByISIN(operation.ISIN)
			if err != nil {
				s.logger.Errorf("Error get stock deal id by isin from operation id:%d, with error: %+v", i, err)
				continue
			}

			if stockdealID == 0 {
				stockInstrument, err := s.repository.StockInstrument().GetInstrumentByISIN(operation.ISIN)
				if err != nil {
					s.logger.Errorf("Error get stock instrument by isin from operation id:%d, with error: %+v", i, err)
					continue
				}

				stockDeal := &models.StockDeal{
					Stock:         *stockInstrument,
					Currency:      &operation.Currency,
					EnterDateTime: operation.DateTime,
					EnterPoint:    operation.Price,
					Quantity:      operation.Quantity,
					UserID:        s.session.userId,
					Variability:   false,
				}

				*stock_deals = append(*stock_deals, *stockDeal)

				continue
			}

		}
	}
}
