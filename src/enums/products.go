package enums

import (
	"database/sql/driver"
	"errors"
)

type ProductStatus string

func (s *ProductStatus) Scan(value string) error {
	*s = ProductStatus(value)
	return nil
}

func (s ProductStatus) Value() (driver.Value, error) {
	return string(s), nil
}

var ProductStatuses = struct {
	Active         ProductStatus
	Inactive       ProductStatus
	OutOfStock     ProductStatus
	Pending        ProductStatus
	Discontinued   ProductStatus
	PreOrder       ProductStatus
	BackOrder      ProductStatus
	ComingSoon     ProductStatus
	LimitedEdition ProductStatus
	OnSale         ProductStatus
}{
	Active:         "active",
	Inactive:       "inactive",
	OutOfStock:     "out-of-stock",
	Pending:        "pending",
	Discontinued:   "discontinued",
	PreOrder:       "pre-order",
	BackOrder:      "back-order",
	ComingSoon:     "coming-soon",
	LimitedEdition: "limited-edition",
	OnSale:         "on-sale",
}

func (s ProductStatus) IsValid() error {
	switch s {
	case ProductStatuses.Active,
		ProductStatuses.Inactive,
		ProductStatuses.OutOfStock,
		ProductStatuses.Pending,
		ProductStatuses.Discontinued,
		ProductStatuses.PreOrder,
		ProductStatuses.BackOrder,
		ProductStatuses.ComingSoon,
		ProductStatuses.LimitedEdition,
		ProductStatuses.OnSale:
		return nil
	}
	return errors.New("invalid ProductStatus")
}
