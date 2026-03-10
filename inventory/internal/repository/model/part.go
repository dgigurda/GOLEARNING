package model

import "time"

type Category int32

const (
	CategoryUnspecified Category = iota
	CategoryEngine
	CategoryFuel
	CategoryPorthole
	CategoryWing
)

type Part struct {
	ID            int64
	UUID          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    *Dimensions
	Manufacturer  *Manufacturer
	Tags          []string
	Metadata      map[string]Metadata
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Manufacturer struct {
	ID      int64
	Name    string
	Country string
	Website string
}

type Dimensions struct {
	ID     int64
	Length float64
	Width  float64
	Height float64
	Weight float64
}

type Metadata struct {
	ID    int64
	Value MetadataValue
}

type MetadataValue interface {
	isMetadataValue()
}

type StringValue struct{ V string }
type Int64Value struct{ V int64 }
type DoubleValue struct{ V float64 }
type BoolValue struct{ V bool }

func (StringValue) isMetadataValue() {}
func (Int64Value) isMetadataValue()  {}
func (DoubleValue) isMetadataValue() {}
func (BoolValue) isMetadataValue()   {}

type GetPartFilter struct {
	ID   int64
	UUID string
}

type ListPartsFilter struct {
	UUIDs                 []string
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}
