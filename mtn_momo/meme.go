package main

import (
	"fmt"
	"margpak/mathformula"
)

func main() {

	var option int
	fmt.Println("Menu Options")
	fmt.Println("1) Transfer Money")
	fmt.Println("2) MoMoPay & Pay Bill")
	fmt.Println("3) Airtime & Bundles")
	fmt.Println("4) Allow Cash Out ")
	fmt.Println("5) Finanacial Services")
	fmt.Println("6) My Wallet")
	fmt.Println("7) MoMo Promo")
	fmt.Println("0) Exit")

	fmt.Print("Enter option: ")
	_, err := fmt.Scanln(&option)
	if err != nil {
		fmt.Println(err)

	}

	// var payment float64
	// var reference string
	// var number int

	switch option {

	case 1:
		// transfer money options
		var payment float64
		var reference string
		var number int
		var transferMoney int

		fmt.Println("Transfer money")
		fmt.Println("1) MoMo User")
		fmt.Println("2) Non MoMo User")
		fmt.Println("3) Send with Care")
		fmt.Println("4) Favorite")
		fmt.Println("5) Other Networks")
		fmt.Println("6) Bank Account")
		fmt.Println("0) Back")

		fmt.Scanln(&transferMoney)
		// _, err := fmt.Scanln(&transferMoney)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		//}

		if transferMoney == 1 { //momo user
			fmt.Println("Enter mobile number (10 digits): ")
			_, err := fmt.Scanln(&number)
			if err != nil || len(fmt.Sprint(number)) != 10 {
				fmt.Println("Invalid mobile number")
				return
			}
			if number != 10 {
				fmt.Errorf("Error, try again.")
				//return
			}

			fmt.Print("Confirm number: ")
			fmt.Scan(&number)

			fmt.Print("Enter amount: ")
			_, err = fmt.Scan(&payment)
			if err != nil {
				fmt.Println("Invalid amount.")
				//return
			}

			currentBalance, err := mathformula.CurrentBalance(3000.10, payment)
			if err != nil {
				fmt.Println(err)
				//	return

			}

			fmt.Print("Enter reference: ")
			fmt.Scanln(&reference)

			//fmt.Println("Your current balance is:", currentBalance)

			fmt.Println("Payment for GHS", payment, "to", number, "Current Balance: ", currentBalance, "Reference:", reference)

		}
		return

		//} else if transferMoney == 2 {

	}

	// 	//case 2:
	// 		fmt.Println("MoMoPay & Pay Bill")
	// 		fmt.Println("1) MoMoPay")
	// 		fmt.Println("2) Pay Bill")
	// 		fmt.Println("3) GhQR")
	// 		fmt.Println("0) Back")

	// 	case 3:
	// 		fmt.Println("Airtime & Bundles")
	// 		fmt.Println("1) Airtime")
	// 		fmt.Println("2) Internet Bundles")
	// 		fmt.Println("3) Fixed Broadband")
	// 		fmt.Println("4) Schedule Airtime")
	// 		fmt.Println("5) Just4U")
	// 		fmt.Println("0) Back")

	// 	case 4:
	// 		fmt.Println("Allow Cash Out")
	// 		fmt.Println("1) Yes")
	// 		fmt.Println("2) No")
	// 		fmt.Println("0) Back")

	// 	case 5:
	// 		fmt.Println("Finanacial Services")
	// 		fmt.Println("1) Bank Services")
	// 		fmt.Println("2) Savings")
	// 		fmt.Println("3) Loans")
	// 		fmt.Println("4) Pensions and Investments")
	// 		fmt.Println("5) Insurance")
	// 		fmt.Println("6) Trade Shares")
	// 		fmt.Println("0) Back")

	// 	case 6:
	// 		fmt.Println("My Wallet")
	// 		fmt.Println("1) MoMo User")
	// 		fmt.Println("2) Non MoMo User")
	// 		fmt.Println("3) Send with Care")
	// 		fmt.Println("4) Favorite")
	// 		fmt.Println("5) Other Networks")
	// 		fmt.Println("6) Bank Account")
	// 		fmt.Println("0) Back")

	// 	case 7:
	// 		fmt.Println("MoMo Promo")
	// 		fmt.Println("1) Check Promo Points")
	// 		fmt.Println("0) Back")
	// 		var MomoPromo int

	// 		fmt.Print("Enter option: ")
	// 		_, err := fmt.Scanln(&MomoPromo)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}

	// if MomoPromo == 1 {
	// 	fmt.Print("Yello, you have 0 points in the MoMo Promo. Use the MoMO App wallet to make payments to Merchant IDs and QR Codes to earn points")

	//return
}
