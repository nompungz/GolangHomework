package nompungz

var roman = []string{"C","XC","L","XL","X","IX","V","IV","I"}
var number = []int{100,90,50,40,10,9,5,4,1}


func romanConverter(n int)string {
	romanString := ""
	for i:= 0 ; i <len(number); i++ {
		for number[i] <= n {
			romanString += roman[i]
			n -= number[i]
		}
	}
	return  romanString
}