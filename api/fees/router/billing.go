package router

import (
	"context"

	"github.com/dsha256/sharingan/api/fees/request"
	"github.com/dsha256/sharingan/api/fees/response"
)

//encore:api public method=POST path=/api/fees/v1/bills
func (s *Fees) OpenBill(ctx context.Context, req *request.OpenBill) (*response.OpenBill, error) {
	return s.billingController.OpenBill(ctx, req)
}

//encore:api public method=GET path=/api/fees/v1/bills/:id
func (s *Fees) GetBill(ctx context.Context, id string) (*response.GetBill, error) {
	return s.billingController.GetBill(ctx, id)
}

//encore:api public method=GET path=/api/fees/v1/bills
func (s *Fees) ListBills(ctx context.Context, query *request.ListBills) (*response.ListBills, error) {
	return s.billingController.ListBills(ctx, query)
}

//encore:api public method=POST path=/api/fees/v1/bills/active/:id/items
func (s *Fees) AddItem(ctx context.Context, id string, req *request.AddItem) (*response.AddItem, error) {
	return s.billingController.AddItem(ctx, id, req)
}

//encore:api public method=POST path=/api/fees/v1/bills/active/:id/close
func (s *Fees) CloseBill(ctx context.Context, id string) (*response.CloseBill, error) {
	return s.billingController.CloseBill(ctx, id)
}
