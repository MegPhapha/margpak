package main

import (
	"fmt"
	"math/rand"
)

type Profile struct {
	MobileNumber string
	Fullname     string
	DateOfBirth  string
}

var profiles = []Profile{
	{MobileNumber: "0111111111", Fullname: "Kofi Manu", DateOfBirth: "31/09/1995"},
	{MobileNumber: "0222222222", Fullname: "Meg Phapha", DateOfBirth: "05/12/2003"},
	{MobileNumber: "0333333333", Fullname: "Addisson Jade", DateOfBirth: "09/07/1993"},
	{MobileNumber: "0444444444", Fullname: "John Mensah", DateOfBirth: "05/01/2005"},
	{MobileNumber: "0555555555", Fullname: "Gina Nun", DateOfBirth: "06/04/2006"},
	{MobileNumber: "0666666666", Fullname: "Addo Dankwa", DateOfBirth: "09/08/1990"},
	{MobileNumber: "0777777777", Fullname: "Mama Cee", DateOfBirth: "05/02/2000"},
}

func main() {
	// Variable declarations
	var options int
	var number string
	var confirmed string
	var payment float64
	var availableBalance float64 = 4000.00
	var reference string
	var recipientName string
	var pinCode int

	fmt.Println("Input short code (*170#)")
	fmt.Println("1) Transfer Money")
	fmt.Println("2) MoMoPay & Pay Bill")
	fmt.Println("3) Airtime & Bundles")
	fmt.Println("4) Allow Cash Out")
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

			// Confirm name of recipient
			for _, profile := range profiles {
				if profile.MobileNumber == number {
					recipientName = profile.Fullname
					break
				}
			}
			if recipientName == "" {
				fmt.Println("Recipient name not found. Please check the number and try again")
				return
			}
			fmt.Println("Recipient:", recipientName)
		}
		

		fmt.Println("Enter amount: ")
		fmt.Scan(&payment)
		if err != nil || payment <= 0 {
			fmt.Println("Invalid amount.")
			return
		}

		fmt.Println("Enter reference: ")
		fmt.Scan(&reference)
		
		// calculate fee charged
		fee := payment * 0.015

		// Calculate e-levy amount(1% of payment amount)
		levy := payment * 0.01

		// Check if payment + e-levy exceeds available balance

		if payment+ fee +levy > availableBalance {
			fmt.Println("Your balance is insufficient")
			return
		}
		

		// Deduct payment + e-levy from available balance to know current balance
		currentBalance := availableBalance - payment -fee - levy

		total := payment + fee + levy 
		fmt.Println("Transfer to", recipientName,"[",number,"]","for GHS",payment,
		 "with Reference:",reference,". Fee is GHS",fee, "Tax amount is GHS",levy,". Total amount is GHS",total)


		 // Enter pincode to confirm/approve payment
		 fmt.Println("Enter pin code")
		 fmt.Scan(&pinCode)
		 if pinCode != 9393{ 
		 fmt.Println("Wrong pin, please try again")
		 return}
		


		// Generate a unique transaction Id
		transctionId := rand.Intn(1000000000055555)

		fmt.Println("Payment made for GHS",payment, "to", recipientName,"[",number,"]", "Current Balance:",currentBalance, ".",
		 "Reference:",reference,".", "Transaction ID:", transctionId,"Fee charged: GHS",fee, "Tax charged GHS:",levy)

		



		//} else if options == 2 {

		return

	case 2:
		//	MoMoPay & Pay Bill options
		fmt.Println("MoMoPay & Pay Bill")
		fmt.Println("1) MoMoPay")
		fmt.Println("2) Pay Bill")
		fmt.Println("3) GhQR")
		fmt.Println("0) Back")

	case 3:
		// Airtime and bundles options
		fmt.Println("Airtime & Bundles")
		fmt.Println("1) Airtime")
		fmt.Println("2) Internet Bundles")
		fmt.Println("3) Fixed Broadband")
		fmt.Println("4) Schedule Airtime")
		fmt.Println("5) Just4U")
		fmt.Println("0) Back")

	case 4:
		// Allow Cash Out options
		fmt.Println("Allow Cash Out")
		fmt.Println("1) Yes")
		fmt.Println("2) No")
		fmt.Println("0) Back")

	case 5:
		// Financial Services options
		fmt.Println("Finanacial Services")
		fmt.Println("1) Bank Services")
		fmt.Println("2) Savings")
		fmt.Println("3) Loans")
		fmt.Println("4) Pensions and Investments")
		fmt.Println("5) Insurance")
		fmt.Println("6) Trade Shares")
		fmt.Println("0) Back")

	case 6:
		//My Wallet options
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
