package types

import "github.com/Orololuwa/collect_am-api/src/dtos"

type EditCustomerPayload struct {
	dtos.FindByID
	dtos.UpdateCustomer
}

type GetACustomerPayload struct {
	dtos.FindByID
}
