package types

import (
	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/repository"
)

type GetAllProductsPayload struct {
	repository.PaginationFilter
	dtos.GetAllProductsFilter
}
