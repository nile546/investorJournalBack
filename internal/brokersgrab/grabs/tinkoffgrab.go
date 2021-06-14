package grabs

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/nile546/diplom/internal/models"
)

type TinkoffGrab struct {
	token string
}

type extoStock struct {
	ID              int64      `json:"id"`
	EnterDateTime   time.Time  `json:"enter_datetime"`
	EnterPoint      int64      `json:"enter_point"`
	StopLoss        *int64     `json:"stop_loss"`
	Quantity        int        `json:"quantity"`
	ExitDateTime    *time.Time `json:"exit_datetime"`
	ExitPoint       *int64     `json:"exit_point"`
	RiskRatio       float32    `json:"risk_ratio"`
	Result          *int64     `json:"result"`
	ResultInPercent float64    `json:"result_in_percent"`
	StartDeposit    int64      `json:"start_deposit"`
	EndDeposit      int64      `json:"end_deposit"`
	UserID          int64      `json:"user_id"`
}

func (t *TinkoffGrab) GetTinkoffStockDeals(token string) (*[]models.StockDeal, error) {
	t.token = token
	t.getTinkoffOperations()
	return nil, nil
}

func (t *TinkoffGrab) getTinkoffOperations() {

	client := sdk.NewRestClient(t.token)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	operations, err := client.Operations(ctx, sdk.DefaultAccount, time.Now().AddDate(0, 0, -100), time.Now(), "")
	if err != nil {
		fmt.Println(err)
	}

	//for _, operation := range operations {

	//}

	fmt.Println(operations)

}
