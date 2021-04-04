package stats

import (
	"github.com/Sonicspeedfly/bank/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа .
func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	var i types.Money
	if len(payments) < 1 {
		return 0
	}
	for _, payment := range payments {
		sum += payment.Amount
		i++
	}
	sred := sum / i
	return sred
}

// TotalInCategory находит сумму покупок в определённой категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, payment := range payments {
		if payment.Category == category{
			sum += payment.Amount
		} 
	}
	return sum
}