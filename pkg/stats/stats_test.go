package stats
	
import (
	"reflect"
	"github.com/Sonicspeedfly/bank/v2/pkg/types"
	"testing"
)

func Test(t *testing.T) {
	payments := []types.Payment{
		{Category: "a", Amount: 10_000,},
		{Category: "a", Amount: 15_000,},
		{Category: "b", Amount: 10_000,},
	}
	expected := map[types.Category]types.Money{
		"a" : 12500,
		"b" : 10000,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}