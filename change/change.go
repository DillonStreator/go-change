package change

import (
	"errors"
	"fmt"
	"sort"
)

// Input Input for Change2 method
type Input struct {
	Owed   int   `json:"owed"`
	Paid   int   `json:"paid"`
	Drawer []int `json:"drawer"`
}

// Calculate produces an output slice of change required given an amount paid, amount owed, and a drawer
func Calculate(input Input) ([]int, error) {
	if input.Paid < input.Owed {
		return nil, errors.New("paid less than owed")
	}

	drawer := input.Drawer

	drawerSize := len(drawer)
	if drawerSize == 0 {
		return nil, errors.New("drawer empty")
	}

	if input.Paid == input.Owed {
		return []int{}, nil
	}

	// reverse sort the drawer
	sort.Sort(sort.Reverse(sort.IntSlice(drawer)))

	var change []int

	remainder := input.Paid - input.Owed
	denominationIdx := 0
	denomination := drawer[denominationIdx]
	for remainder > 0 {
		if denomination > remainder {
			denominationIdx++
			if denominationIdx >= drawerSize {
				return nil, fmt.Errorf("unable to produce adequate change with remainder %d", remainder)
			}
			denomination = drawer[denominationIdx]
			continue
		}

		change = append(change, denomination)
		remainder -= denomination
	}

	return change, nil
}
