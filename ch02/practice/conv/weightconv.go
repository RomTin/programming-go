package conv

import "fmt"

type Kilogram float64
type Pound float64

func (k Kilogram) String() string { return fmt.Sprintf("%gKg", k) }
func (p Pound) String() string { return fmt.Sprintf("%glb", p) }
