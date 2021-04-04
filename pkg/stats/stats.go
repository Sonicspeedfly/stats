package stats

import (
	"github.com/Sonicspeedfly/bank/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа .
func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	var i types.Money
	if len(payments) < 1 {
		return 0
	}
	
	for _, payment := range payments {
		if payment.Status != types.StatusFail{
		sum += payment.Amount
		i++
		}
	}
	sred := sum / i
	return sred
}

// TotalInCategory находит сумму покупок в определённой категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, payment := range payments {
		if ((payment.Category == category)&&(payment.Status != types.StatusFail)){
			sum += payment.Amount
		} 
	}
	return sum
}