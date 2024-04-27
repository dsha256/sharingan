package response

import "github.com/dsha256/sharingan/internal/dto"

type (
	BillData struct {
		Bill *dto.Bill `json:"bill"`
	}
	BillsData struct {
		Bills []*dto.Bill `json:"bills"`
	}
	ItemData struct {
		Item *dto.Item `json:"item"`
	}

	OpenBill struct {
		Data BillData `json:"data"`
	}

	GetBill struct {
		Data BillData `json:"data"`
	}

	ListBills struct {
		Data BillsData `json:"data"`
	}

	AddItem struct {
		Data ItemData `json:"data"`
	}

	CloseBill struct {
		Data BillData `json:"data"`
	}
)
