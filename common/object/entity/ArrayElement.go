package entity

// ElementRefType is the type to define a ref to ArrayElement.
type ElementRefType uint64

// Element represents an element in an array.
type Element struct {
	ID uint64 `gorm:"column:ID;primary_key"`
}
