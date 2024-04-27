package controller

import (
	"context"

	"github.com/dsha256/sharingan/api/fees/request"
	"github.com/dsha256/sharingan/api/fees/response"
	"github.com/dsha256/sharingan/api/fees/service"
	"github.com/dsha256/sharingan/internal/dto"
)

type Billing struct {
	billingService *service.Billing
}

func NewBillingController(billingService *service.Billing) *Billing {
	return &Billing{billingService}
}

func (s *Billing) OpenBill(ctx context.Context, req *request.OpenBill) (*response.OpenBill, error) {
	bill, err := s.billingService.OpenBill(ctx, &dto.Bill{Currency: req.Currency})
	if err != nil {
		return &response.OpenBill{}, response.Error(err)
	}

	var resp response.OpenBill
	resp.Data.Bill = bill.Denilify()

	return &resp, nil
}

func (s *Billing) GetBill(ctx context.Context, id string) (*response.GetBill, error) {
	bill, err := s.billingService.GetBillById(ctx, id)
	if err != nil {
		return &response.GetBill{}, response.Error(err)
	}

	var resp response.GetBill
	resp.Data.Bill = bill.Denilify()

	return &resp, nil
}

func (s *Billing) ListBills(ctx context.Context, query *request.ListBills) (*response.ListBills, error) {
	bills, err := s.billingService.ListBills(ctx, dto.BillStatus(query.Status))
	if err != nil {
		return &response.ListBills{}, response.Error(err)
	}

	for _, bill := range bills {
		bill = bill.Denilify()
	}

	var resp response.ListBills
	resp.Data.Bills = bills

	return &resp, nil
}

func (s *Billing) AddItem(ctx context.Context, id string, req *request.AddItem) (*response.AddItem, error) {
	item, err := s.billingService.AddItem(ctx, id, &dto.Item{
		Name:     req.Name,
		Price:    req.Price,
		Quantity: req.Quantity,
	})
	if err != nil {
		return &response.AddItem{}, response.Error(err)
	}

	var resp response.AddItem
	resp.Data.Item = item

	return &resp, nil
}

func (s *Billing) CloseBill(ctx context.Context, id string) (*response.CloseBill, error) {
	bill, err := s.billingService.CloseBill(ctx, id)

	if err != nil {
		return &response.CloseBill{}, response.Error(err)
	}

	var resp response.CloseBill
	resp.Data.Bill = bill.Denilify()

	return &resp, nil
}
