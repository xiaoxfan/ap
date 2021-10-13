package ap

// Int64DualPrices are dual prices for the assignment constraints corresponding
// to the U and V sets, respectively.
type Int64DualPrices struct {
	U []int64 `json:"u"`
	V []int64 `json:"v"`
}

// An Int64DualPricer provides dual prices on the assignment constraints
// associated with sets U and V. A dual price is the value of a unit of slack on
// a binding constraint.
type Int64DualPricer interface {
	DualPrices() Int64DualPrices
}

// Int64DualPrices are dual prices for the assignment constraints corresponding
// to the U and V sets, respectively.
type Float64DualPrices struct {
	U []float64 `json:"u"`
	V []float64 `json:"v"`
}

// An Int64DualPricer provides dual prices on the assignment constraints
// associated with sets U and V. A dual price is the value of a unit of slack on
// a binding constraint.
type Float64DualPricer interface {
	DualPrices() Float64DualPrices
}
