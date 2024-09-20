package enums

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type IDiscountType string

var EDiscountType = struct {
	Fixed      IDiscountType
	Percentage IDiscountType
}{
	Fixed:      "fixed",
	Percentage: "percentage",
}

func (d *IDiscountType) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*d = IDiscountType(v)
	case string:
		*d = IDiscountType(v)
	default:
		return errors.New("invalid type for DiscountType")
	}
	return nil
}

func (d IDiscountType) Value() (driver.Value, error) {
	switch d {
	case EDiscountType.Fixed, EDiscountType.Percentage:
		return string(d), nil
	}
	return nil, fmt.Errorf("invalid DiscountType: %s", d)
}

type IInvoiceStatus string

var EInvoiceStatus = struct {
	Draft     IInvoiceStatus
	Pending   IInvoiceStatus
	Sent      IInvoiceStatus
	Viewed    IInvoiceStatus
	Paid      IInvoiceStatus
	Overdue   IInvoiceStatus
	Cancelled IInvoiceStatus
	Refunded  IInvoiceStatus
	Failed    IInvoiceStatus
}{
	Draft:     "draft",
	Pending:   "pending",
	Sent:      "sent",
	Viewed:    "viewed",
	Paid:      "paid",
	Overdue:   "overdue",
	Cancelled: "cancelled",
	Refunded:  "refunded",
	Failed:    "failed",
}

func (s *IInvoiceStatus) Scan(value string) error {
	*s = IInvoiceStatus(value)
	return nil
}

func (s IInvoiceStatus) Value() (driver.Value, error) {
	return string(s), nil
}
