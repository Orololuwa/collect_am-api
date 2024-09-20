package enums

import "database/sql/driver"

type ICustomerType string

var ECustomerType = struct {
	Individual ICustomerType
	Corporate  ICustomerType
}{
	Individual: "individual",
	Corporate:  "corporate",
}

func (s *ICustomerType) Scan(value string) error {
	*s = ICustomerType(value)
	return nil
}

func (s ICustomerType) Value() (driver.Value, error) {
	return string(s), nil
}
