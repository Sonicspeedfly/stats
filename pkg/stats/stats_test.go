package stats
	
import (
	"reflect"
	"github.com/Sonicspeedfly/bank/v2/pkg/types"
	"testing"
)

func TestCategoryAvg(t *testing.T) {
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

func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
		"a" : 10,
		"b" : 20,
	}
	second := map[types.Category]types.Money{
		"a" : 5,
		"b" : 3,
	}
	expected := map[types.Category]types.Money{
		"a" : -5,
		"b" : -17,
	}

	result := PeriodsDynamic(first,second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}