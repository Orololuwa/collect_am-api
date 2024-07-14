package dtos

type FindByID struct {
	Id uint `json:"id" faker:"oneof:2,3,4,5"`
}
