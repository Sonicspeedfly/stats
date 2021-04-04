package stats

import (
	"github.com/Sonicspeedfly/bank/pkg/types"
	"fmt"
)

func ExamleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount: 10_000,
			Category: "a",
		},
		{
			Amount: 15_000,
			Category: "a",
		},
		{
			Amount: 10_000,
			Category: "b",
		},
	}
	sum := TotalInCategory(payments,"a")
	fmt.Println(sum)

	//Output:
	//25000
}