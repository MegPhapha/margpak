package main

import (
	"fmt"
	"margpak/mathformula"
)

func main() {

	var options int
	var number int
	var payment float64
	var reference string
	fmt.Println("Input short code (*170#)")
	fmt.Println("1) Transfer Money")
	fmt.Println("2) MoMoPay & Pay Bill")
	fmt.Println("3) Airtime & Bundles")
	fmt.Println("4) Allow Cash Out ")
	fmt.Println("5) Finanacial Services")
	fmt.Println("6) My Wallet")
	fmt.Println("7) MoMo Promo")

	fmt.Print("Enter options: ")
	fmt.Scanln(&options)

	switch options {

	case 1:
		fmt.Println("Transfer Money")
		fmt.Println("1) MoMo User")
		fmt.Println("2) Non MoMo User")
		fmt.Println("3) Send with Care")
		fmt.Println("4) Favorite")
		fmt.Println("5) Other Networks")
		fmt.Println("6) Bank Account")
		fmt.Println("0) Back")

		var TransferMoney int

		fmt.Print("Enter option: ")
		_, err := fmt.Scanln(&TransferMoney)
		if err != nil {
			fmt.Println(err)
			return
		}
		// transfering money to momo user
		if TransferMoney == 1 {

			fmt.Print("Enter mobile number(10 digits): ")
			_, err := fmt.Scanln(&number)
			if err != nil || len(fmt.Sprint(&number)) != 10 {
				fmt.Println("Invalid mobile number.")

				fmt.Println("Confirm number: ")
				fmt.Scan(&number)
			}

			fmt.Println("Enter amount: ")
			fmt.Scan(&payment)

			//formula for current balance in the math formula folder/file
			// func CurrentBalance(availablebal float64, payment float64) (float64, error) {
			// 	if payment > availablebal {
			// 		return 0, fmt.Errorf("insufficient balance")

			// 	}

			// 	current_bal := availablebal - payment

			// 	return current_bal, nil
			currentBalance, err := mathformula.CurrentBalance(3000.10, payment)
			if err != nil || payment <= 0 {
				fmt.Println("Invalid amount.")

				if payment > currentBalance {
					fmt.Println("Insufficient balance.")
				}
			}

			fmt.Println("Enter reference: ")
			fmt.Scan(&reference)
			//initializing transfer
			//currentBalance -= payment
			fmt.Println("Your current balance is:", currentBalance)

			fmt.Println("Payment for GHS", payment, "to", number, "Current Balance: ", currentBalance, "Reference:", reference)
			return

		} else if TransferMoney == 2 {

		}

	case 2:
		fmt.Println("MoMoPay & Pay Bill")
		fmt.Println("1) MoMoPay")
		fmt.Println("2) Pay Bill")
		fmt.Println("3) GhQR")
		fmt.Println("0) Back")

	case 3:
		fmt.Println("Airtime & Bundles")
		fmt.Println("1) Airtime")
		fmt.Println("2) Internet Bundles")
		fmt.Println("3) Fixed Broadband")
		fmt.Println("4) Schedule Airtime")
		fmt.Println("5) Just4U")
		fmt.Println("0) Back")

	case 4:
		fmt.Println("Allow Cash Out")
		fmt.Println("1) Yes")
		fmt.Println("2) No")
		fmt.Println("0) Back")

	case 5:
		fmt.Println("Finanacial Services")
		fmt.Println("1) Bank Services")
		fmt.Println("2) Savings")
		fmt.Println("3) Loans")
		fmt.Println("4) Pensions and Investments")
		fmt.Println("5) Insurance")
		fmt.Println("6) Trade Shares")
		fmt.Println("0) Back")

	case 6:
		fmt.Println("My Wallet")
		fmt.Println("1) MoMo User")
		fmt.Println("2) Non MoMo User")
		fmt.Println("3) Send with Care")
		fmt.Println("4) Favorite")
		fmt.Println("5) Other Networks")
		fmt.Println("6) Bank Account")
		fmt.Println("0) Back")

	case 7:
		fmt.Println("MoMo Promo")
		fmt.Println("1) Check Promo Points")
		fmt.Println("0) Back")
		var MomoPromo int

		fmt.Print("Enter option: ")
		_, err := fmt.Scanln(&MomoPromo)
		if err != nil {
			fmt.Println(err)
			return
		}

		if MomoPromo == 1 {
			fmt.Print("Yello, you have 0 points in the MoMo Promo. Use the MoMO App wallet to make payments to Merchant IDs and QR Codes to earn points")

			return
		}
	}
}