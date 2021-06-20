package apiserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nile546/diplom/internal/models"
)

func (s *server) getAllStockDealFromBrokers(w http.ResponseWriter, r *http.Request) {

	type request struct {
		TinkoffToken  string `json:"tinkoff_token"`
		AutoGrabDeals bool   `json:"auto_grab_deals"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	err := s.getTinkoffStockDeals(req.TinkoffToken, req.AutoGrabDeals)
	if err != nil {
		s.logger.Errorf("Error insert Tinkoff stock deals, with error: %+v", err)
		s.error(w, "Request not processed, please try again later, "+err.Error())
	}

	s.respond(w, nil)

}

func (s *server) getTinkoffStockDeals(token string, autoGrabDeals bool) error {

	if !autoGrabDeals {
		err := s.insertTinkoffStockDeals(token)
		if err != nil {
			return err
		}

		err = s.repository.TinkoffToken().InsertTinkoffToken(token, s.session.userId)
		if err != nil {
			s.logger.Errorf("Error update auto grab stock deal, with error: %+v", err)
		}
		return nil
	}

	err := s.repository.User().UpdateAutoGrab(s.session.userId, true)
	if err != nil {
		s.logger.Errorf("Error update auto grab stock deal, with error: %+v", err)
	}

	go func() error {
		for {
			err = s.insertTinkoffStockDeals(token)
			if err != nil {
				s.logger.Errorf("Error insert Tinkoff stock deals, with error: %+v", err)
				break
			}
			time.Sleep(time.Second * 30)
		}
		if err != nil {
			return err
		}
		return nil
	}()
	return nil
}

func (s *server) insertTinkoffStockDeals(token string) error {

	grabDate, err := s.repository.User().GetDateGrabByUserID(s.session.userId)
	if err != nil {
		s.logger.Errorf("Error get date grab stock deal, with error: %+v", err)
		return err
	}

	tinkoff_operations, err := s.brokersGrab.TinkoffGrab().GetTinkoffStockDeals(token, grabDate)
	if err != nil {
		return err
	}

	if tinkoff_operations == nil {
		return nil
	}

	err = s.repository.User().UpdateDateGrab(time.Now(), s.session.userId)
	if err != nil {
		s.logger.Errorf("Error update date grab stock deal, with error: %+v", err)
		return err
	}

	for i, operation := range *tinkoff_operations {
		if operation.Operation == 1 {

			stockDealID := s.repository.StockDeal().GetStockDealsIDByISIN(operation.ISIN)

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
					Variability:   false,
				}

				idStockDeal, err := s.repository.StockDeal().CreateOpenStockDeal(stockDeal, s.session.userId)
				if err != nil {
					s.logger.Errorf("Error create stock deal from operation id:%d, with error: %+v", i, err)
					continue
				}

				stockDealPart := &models.StockDealParts{
					Quantity:    operation.Quantity,
					DealType:    operation.Operation,
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
				DealType:    operation.Operation,
				Price:       operation.Price,
				DateTime:    operation.DateTime,
				StockDealId: stockDealID,
			}

			err := s.repository.StockDealPart().InsertStockDealPart(stockDealPart)
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

		stockDealPart := &models.StockDealParts{
			Quantity:    operation.Quantity,
			DealType:    operation.Operation,
			Price:       operation.Price,
			DateTime:    operation.DateTime,
			StockDealId: stockDealID,
		}

		err := s.repository.StockDealPart().InsertStockDealPart(stockDealPart)
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
		if err != nil {
			s.logger.Errorf("Error set stock deal completed in parts from operation id:%d, with error: %+v", i, err)
		}
	}
	return nil
}

func (s *server) offAutoGrabDeals() {

}
