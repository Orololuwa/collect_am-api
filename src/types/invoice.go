package types

import "github.com/Orololuwa/collect_am-api/src/dtos"

type CreateInvoicePayload struct {
	dtos.CreateInvoice
}

type GetAnInvoicePayload struct {
	dtos.FindByID
}
