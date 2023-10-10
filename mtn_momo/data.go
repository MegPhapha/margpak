package main

import (
	"fmt"
	"math/rand"
)

// Define a map to store mobile numbers and recipient names
var recipientDatabase = map[string]string{
	"0111111111": "Kofi Manu",
	"0222222222": "Meg Phapha",
	"0333333333": "Addisson Jade",
	"0444444444": "John Mensah",
	"0555555555": "Gina Nun",
	"0666666666": "Addo Dankwa",
	"0777777777": "Mama Cee",
}

func main() {
	// Variable declarations
	var options int
	var number string
	var confirmed string
	var payment float64
	var availableBalance float64 = 4000.00
	var reference string

	fmt.Println("Input short code (*170#)")
	fmt.Println("1) Transfer Money")
	fmt.Println("2) MoMoPay & Pay Bill")
	fmt.Println("3) Airtime & Bundles")
	fmt.Println("4) Allow Cash Out ")
	fmt.Println("5) Financial Services")
	fmt.Println("6) My Wallet")
	fmt.Println("7) MoMo Promo")

	fmt.Print("Enter options: ")
	fmt.Scan(&options)

	switch options {
	case 1:
		// Transfer Money option
		fmt.Println("Transfer Money")
		fmt.Println("1) MoMo User")
		fmt.Println("2) Non MoMo User")
		fmt.Println("3) Send with Care")
		fmt.Println("4) Favorite")
		fmt.Println("5) Other Networks")
		fmt.Println("6) Bank Account")
		fmt.Println("0) Back")

		fmt.Print("Enter option: ")
		_, err := fmt.Scan(&options)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Transfer money to MoMo user
		if options == 1 {
			fmt.Print("Enter mobile number (10 digits): ")
			_, err := fmt.Scan(&number)
			if err != nil || len(number) != 10 {
				fmt.Println("Invalid mobile number.")
				return
			}

			// Check if the recipient's name is in the database
			recipientName, found := recipientDatabase[number]
			if !found {
				fmt.Println("Recipient not found.")
				return
			}

			fmt.Printf("Recipient: %s\n", recipientName)

			fmt.Print("Confirm number: ")
			_, err = fmt.Scan(&confirmed)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Check if confirmed number matches the entered mobile number
			if number != confirmed {
				fmt.Println("Number mismatched, try again.")
				return
			}

			fmt.Println("Enter amount: ")
			fmt.Scan(&payment)
			if err != nil || payment <= 0 {
				fmt.Println("Invalid amount.")
				return
			}

			// Check if payment exceeds available balance
			if payment > availableBalance {
				fmt.Println("Insufficient balance")
				return
			}

			// Deduct payment from available balance to know the current balance
			currentBalance := availableBalance - payment
			fmt.Println("Your current balance is:", currentBalance)

			fmt.Println("Enter reference: ")
			fmt.Scan(&reference)

			// Generate a unique transaction ID
			transactionID := rand.Intn(1000000000055555)

			fmt.Printf("Payment made for GHS %.2f to %s. Current Balance: %.2f. Reference: %s. Transaction ID: %d\n",
				payment, recipientName, currentBalance, reference, transactionID)

			return
		}

		// Handle other transfer options...

		// Return after processing the transfer
		return
	}

	// Handle other menu options...

}
