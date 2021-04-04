package stats

import (
	"github.com/Sonicspeedfly/bank/v2/pkg/types"
	"fmt"
)

func ExamleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount: 10_000,
			Category: "a",
			Status: "OK",
		},
		{
			Amount: 15_000,
			Category: "a",
			Status: "INPROGRESS",
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