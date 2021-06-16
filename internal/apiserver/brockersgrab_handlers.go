package apiserver

import (
	"encoding/json"
	"net/http"
	"time"

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

	grabDate, err := s.repository.User().GetDateGrabByUserID(1)
	if err != nil {
		s.logger.Errorf("Error get date grab stock deal, with error: %+v", err)
	}

	tinkoff_operations, err := s.brokersGrab.TinkoffGrab().GetTinkoffStockDeals(req.token, grabDate)
	if err != nil {
		s.respond(w, err.Error())
		return
	}

	err = s.repository.User().UpdateDateGrab(time.Now(), 1)
	if err != nil {
		s.logger.Errorf("Error update date grab stock deal, with error: %+v", err)
	}

	for i, operation := range *tinkoff_operations {
		if operation.Operation == 1 {

			stockDealID := s.repository.StockDeal().GetStockDealsIDByISIN(operation.ISIN)
			//if err != nil {
			//	s.logger.Errorf("Error get stock deal id by isin from operation id:%d, with error: %+v", i, err)
			//	continue
			//}

			if stockDealID == 0 {
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

				idStockDeal, err := s.repository.StockDeal().CreateOpenStockDeal(stockDeal)
				if err != nil {
					s.logger.Errorf("Error create stock deal from operation id:%d, with error: %+v", i, err)
					continue
				}

				stockDealPart := &models.StockDealParts{
					Quantity:    operation.Quantity,
					Type:        operation.Operation,
					Price:       operation.Price,
					DateTime:    operation.DateTime,
					StockDealId: idStockDeal,
				}

				err = s.repository.StockDealPart().InsertStockDealPart(stockDealPart)
				if err != nil {
					s.logger.Errorf("Error insert stock deal parts from operation id:%d, with error: %+v", i, err)
					continue
				}

				continue
			}

			stockDealPart := &models.StockDealParts{
				Quantity:    operation.Quantity,
				Type:        operation.Operation,
				Price:       operation.Price,
				DateTime:    operation.DateTime,
				StockDealId: stockDealID,
			}

			err = s.repository.StockDealPart().InsertStockDealPart(stockDealPart)
			if err != nil {
				s.logger.Errorf("Error insert stock deal parts from operation id:%d, with error: %+v", i, err)
				continue
			}

			err = s.repository.StockDeal().UpdateQuantityStockDeal(stockDealID, operation.Quantity)
			if err != nil {
				s.logger.Errorf("Error update stock deal parts from operation id:%d, with error: %+v", i, err)
				continue
			}

			continue
		}

		stockDealID := s.repository.StockDeal().GetStockDealsIDByISIN(operation.ISIN)
		//if err != nil {
		//	s.logger.Errorf("Error get stock deal id by isin from operation id:%d, with error: %+v", i, err)
		//	continue
		//}

		stockDealPart := &models.StockDealParts{
			Quantity:    operation.Quantity,
			Type:        operation.Operation,
			Price:       operation.Price,
			DateTime:    operation.DateTime,
			StockDealId: stockDealID,
		}

		err = s.repository.StockDealPart().InsertStockDealPart(stockDealPart)
		if err != nil {
			s.logger.Errorf("Error insert stock deal parts from operation id:%d, with error: %+v", i, err)
			continue
		}

		statusCompletedDeal, err := s.repository.StockDealPart().CheckQuantityDeal(stockDealID)
		if err != nil {
			s.logger.Errorf("Error check stock deal completed in parts from operation id:%d, with error: %+v", i, err)
			continue
		}

		if !statusCompletedDeal {
			continue
		}

		err = s.repository.StockDeal().SetStockDealCompleted(operation.DateTime, operation.Price, stockDealID)

	}
}
