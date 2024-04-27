package workflow

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	"github.com/dsha256/sharingan/internal/dto"
	"github.com/dsha256/sharingan/internal/temporal/activity"
)

const (
	Name     = "Billing"
	IdPrefix = "billing-"

	UpdateOpenBill = "OPEN_BILL"
	UpdateAddItem  = "ADD_ITEM"

	QueryGetBill = "GET_BILL"

	SignalDone = "DONE"
)

var (
	CommonActivityOptions = workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}
)

type (
	WFBillingState = struct {
		Bill dto.Bill
	}
)

func Billing(ctx workflow.Context) error {
	state := WFBillingState{
		Bill: dto.Bill{},
	}

	if err := workflow.SetUpdateHandler(ctx,
		UpdateOpenBill,
		func(ctx workflow.Context, bill *dto.Bill) (*dto.Bill, error) {
			state.Bill = *bill
			state.Bill.Total = "0.0"
			return &state.Bill, nil
		}); err != nil {
		return err
	}

	if err := workflow.SetUpdateHandler(ctx,
		UpdateAddItem,
		func(ctx workflow.Context, item *dto.Item) (*dto.Item, error) {
			var total string
			opts := activity.MoneyAdditionOptions{Target: state.Bill.Total, Addition: item.Price, Times: item.Quantity}
			err := workflow.ExecuteActivity(workflow.WithActivityOptions(ctx, CommonActivityOptions), activity.MoneyAddition, opts).
				Get(ctx, &total)
			if err != nil {
				return nil, err
			}
			state.Bill.Total = total

			state.Bill.Items = append(state.Bill.Items, item)

			return item, nil
		}); err != nil {
		return err
	}

	if err := workflow.SetQueryHandler(ctx,
		QueryGetBill,
		func() (*dto.Bill, error) {
			return &state.Bill, nil
		}); err != nil {
		return err
	}

	_ = workflow.GetSignalChannel(ctx, SignalDone).Receive(ctx, nil)
	state.Bill.ClosedAt = time.Now()

	return nil
}
