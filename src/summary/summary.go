package summary

import (
	"credit-calc/config"
	"credit-calc/transactions"
	"fmt"
	"time"
)

type Summary struct {
	TotalAmountSpent  float64
	TotalPointsEarned int
	AmountLeft        float64
	DaysLeft          int
}

func GenerateSummary(transactions []*transactions.Transaction, config *config.Config, now time.Time) (Summary, error) {
	if config == nil {
		return Summary{
			TotalAmountSpent:  0,
			TotalPointsEarned: 0,
			AmountLeft:        0,
			DaysLeft:          0,
		}, ErrEmptyConfig(nil)
	}

	totalAmountSpent := 0.0
	totalPointsEarned := 0
	amountLeft := config.SpendAmountForBonus
	cardStartDate := config.CardStartDate

	for _, transaction := range transactions {
		if transaction.Date.Before(cardStartDate) {
			continue
		}

		if transaction.IsEligible(config) {
			amountLeft += (-transaction.Amount)
			totalAmountSpent += -transaction.Amount
		}
	}

	totalPointsEarned = int(totalAmountSpent)

	return Summary{
		TotalAmountSpent:  totalAmountSpent,
		TotalPointsEarned: totalPointsEarned,
		AmountLeft:        amountLeft,
		DaysLeft:          config.BonusSpendPeriodInDays - int(now.Sub(cardStartDate).Hours()/24),
	}, nil
}

func (s *Summary) Display() {
	fmt.Printf("Total amount spent: %.2f\n", s.TotalAmountSpent)
	fmt.Printf("Total points earned: %d\n", s.TotalPointsEarned)
	fmt.Printf("Amount left: %.2f\n", s.AmountLeft)
	fmt.Printf("Days left: %d\n", s.DaysLeft)
}
