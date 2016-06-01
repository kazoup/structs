package countvalue

import (
	"github.com/kazoup/structs/intmap"
)

// CountValue data struct
type CountValue struct {
	Count int           `json:"count"`
	Value intmap.Intmap `json:"value"`
}
