package main
import(
	"fmt"

)

type Database struct {
	MobileNumber  string 
	 Fullname string
}

func main () {
	var database = []Database{
		{"0111111111", "Kofi Manu"},
			{"0222222222", "Meg Phapha"},
			{"0333333333",  "Addisson Jade"},
			{"0444444444", "John Mensah"}, 
			{"0555555555", "Gina Nun"},
			{"0666666666", "Addo Dankwa"},
			{"0777777777", "Mama Cee"},

	}

	
fmt.Println(database)
}
