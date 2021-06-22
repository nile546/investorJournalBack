package apiserver

import (
	"github.com/nile546/diplom/internal/models"
)

func (s *server) riskRatio(enterPoint *int64, exitPoint *int64, stopLoss *int64, position *models.Positions) *float64 {
	if enterPoint == nil || exitPoint == nil || stopLoss == nil || position == nil {
		return nil
	}

	var res float64

	if *position == models.Long {
		if float64(*enterPoint-*stopLoss) != 0 {
			res = float64(*exitPoint-*enterPoint) / float64(*enterPoint-*stopLoss)
			return &res
		}
		res = 0

	} else {
		if float64(*stopLoss-*enterPoint) != 0 {
			res = float64(*enterPoint-*exitPoint) / float64(*stopLoss-*enterPoint)
			return &res
		}
		res = 0
	}

	return &res
}

func (s *server) result(enterPoint *int64, exitPoint *int64, quantity *int, position *models.Positions) *int64 {
	if enterPoint == nil || exitPoint == nil || quantity == nil || position == nil {
		return nil
	}

	var res int64

	if *position == models.Long {
		res = (*exitPoint - *enterPoint) * int64(*quantity)
	} else {
		res = (*enterPoint - *exitPoint) * int64(*quantity)
	}

	return &res
}
