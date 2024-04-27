package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.temporal.io/api/filter/v1"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"

	"github.com/dsha256/sharingan/api/fees/config"
	"github.com/dsha256/sharingan/internal/dto"
	"github.com/dsha256/sharingan/internal/temporal/workflow"
)

type Billing struct {
	client client.Client
}

func NewBillingService(temporalClient client.Client) *Billing {
	return &Billing{client: temporalClient}
}

func (s *Billing) OpenBill(ctx context.Context, bill *dto.Bill) (*dto.Bill, error) {
	bill.Id = uuid.New().String()
	bill.CreatedAt = time.Now()

	we, err := s.client.ExecuteWorkflow(ctx,
		client.StartWorkflowOptions{
			ID:        workflow.IdPrefix + bill.Id,
			TaskQueue: config.FeesTaskQueue,
		},
		workflow.Billing,
	)
	if err != nil {
		return &dto.Bill{}, err
	}

	updateHandler, err := s.client.UpdateWorkflow(ctx, we.GetID(), "", workflow.UpdateOpenBill, bill)
	if err != nil {
		return &dto.Bill{}, err
	}

	result := dto.Bill{}
	err = updateHandler.Get(context.Background(), &result)
	if err != nil {
		return &dto.Bill{}, err
	}

	return &result, nil
}

func (s *Billing) GetBillById(ctx context.Context, id string) (*dto.Bill, error) {
	return s.getBillByWorkflowId(ctx, workflow.IdPrefix+id)
}

func (s *Billing) ListBills(ctx context.Context, status dto.BillStatus) (bills []*dto.Bill, e error) {
	var nextPageToken []byte
	var res *dto.Bill

	for {
		switch status {
		case dto.BillingStatusOpen:
			request := &workflowservice.ListOpenWorkflowExecutionsRequest{
				Filters: &workflowservice.ListOpenWorkflowExecutionsRequest_TypeFilter{
					TypeFilter: &filter.WorkflowTypeFilter{
						Name: workflow.Name,
					},
				},
				MaximumPageSize: 10,
				NextPageToken:   nextPageToken,
			}
			response, err := s.client.ListOpenWorkflow(ctx, request)
			if err != nil {
				return
			}

			for _, exec := range response.Executions {
				if res, err = s.getBillByWorkflowId(ctx, exec.Execution.GetWorkflowId()); err != nil {
					return
				}
				bills = append(bills, res)
			}

			nextPageToken = response.NextPageToken
			if len(nextPageToken) == 0 {
				return
			}

		case dto.BillingStatusClosed:
			request := &workflowservice.ListClosedWorkflowExecutionsRequest{
				Filters: &workflowservice.ListClosedWorkflowExecutionsRequest_TypeFilter{
					TypeFilter: &filter.WorkflowTypeFilter{
						Name: workflow.Name,
					},
				},
				MaximumPageSize: 10,
				NextPageToken:   nextPageToken,
			}
			response, err := s.client.ListClosedWorkflow(ctx, request)
			if err != nil {
				return
			}

			for _, exec := range response.Executions {
				if res, err = s.getBillByWorkflowId(ctx, exec.Execution.GetWorkflowId()); err != nil {

					return
				}
				bills = append(bills, res)
			}

			nextPageToken = response.NextPageToken
			if len(nextPageToken) == 0 {
				return
			}
		}
	}
}

func (s *Billing) AddItem(ctx context.Context, billId string, item *dto.Item) (*dto.Item, error) {
	item.CreatedAt = time.Now()

	updateHandler, err := s.client.UpdateWorkflow(ctx, workflow.IdPrefix+billId, "", workflow.UpdateAddItem, item)
	if err != nil {
		return &dto.Item{}, err
	}

	result := dto.Item{}
	err = updateHandler.Get(context.Background(), &result)
	if err != nil {
		return &dto.Item{}, err
	}

	return &result, err
}

func (s *Billing) CloseBill(ctx context.Context, id string) (*dto.Bill, error) {
	if err := s.client.SignalWorkflow(context.Background(), workflow.IdPrefix+id, "", workflow.SignalDone, nil); err != nil {
		return &dto.Bill{}, err
	}

	return s.getBillByWorkflowId(ctx, workflow.IdPrefix+id)
}

func (s *Billing) getBillByWorkflowId(_ context.Context, wfId string) (*dto.Bill, error) {
	queryResp, err := s.client.QueryWorkflow(context.Background(), wfId, "", workflow.QueryGetBill)
	if err != nil {
		return &dto.Bill{}, err
	}

	bill := dto.Bill{}
	if err = queryResp.Get(&bill); err != nil {
		return &dto.Bill{}, err
	}

	return &bill, nil
}
