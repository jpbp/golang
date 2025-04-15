package main
import "fmt"

func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 0.1
	case "pro":
		return 0.2
	case "enterprise":
		return 0.5
	default:
		return 0.0
	}
}

func calculateBalance (accountBalance float64, finalCost float64, planUser string)float64{
	finalCost = finalCost - finalCost * billingCost(planUser)
	return accountBalance - finalCost
}

func validatedSald (accountBalance float64) string{
	
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	if accountBalance >= 0.0 {
		return purchaseSuccessMessage
	}
	return insufficientFundMessage
}

func main() {
	
	var accountBalance float64 = 74.0
	var bulkMessageCost float64 = 75.0
	var planUser string = "free"
	var finalCost float64 = bulkMessageCost

	accountBalance = calculateBalance(accountBalance,finalCost,planUser)

	fmt.Println(validatedSald(accountBalance))
	fmt.Println("Account balance:", accountBalance)
}