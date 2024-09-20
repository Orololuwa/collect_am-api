package enums

import "database/sql/driver"

type IStatus string

var EStatus = struct {
	Active  IStatus
	Blocked IStatus
	Deleted IStatus
}{
	Active:  "active",
	Blocked: "blocked",
	Deleted: "deleted",
}

func (s *IStatus) Scan(value string) error {
	*s = IStatus(value)
	return nil
}

func (s IStatus) Value() (driver.Value, error) {
	return string(s), nil
}
