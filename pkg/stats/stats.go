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

//CategoriesAvg рассчетать среднюю сумму всех категорий
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money{
	i := map[types.Category]types.Money{}
	sum := map[types.Category]types.Money{}
	categories := map[types.Category]types.Money{}
	for _, payment := range payments {
		sum[payment.Category] += payment.Amount
		i[payment.Category]++
		categories[payment.Category] = sum[payment.Category]/i[payment.Category]
	}
	return categories
}