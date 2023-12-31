package models

type CRUDModel[T any] interface {
	Create() T
	GetList() []T
	GetSingleItem(ID int) T
	UpdateItem(ID int) T
	Delete(ID int) bool
	TotalCount() int64
}
