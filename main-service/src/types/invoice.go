package types

import "github.com/Orololuwa/collect_am-api/src/dtos"

type CreateInvoicePayload struct {
	dtos.CreateInvoice
}

type EditInvoicePayload struct {
	ID   uint
	Body map[string]interface{}
}

type EditListedProductPayload struct {
	Id          int
	Quantity    int
	PriceListed float64
}

type GetAnInvoicePayload struct {
	dtos.FindByID
}
