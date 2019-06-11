package nompungz 

import (
	"strings"
)

func phoneMobileNo(inputs []string)map[string]int{
	mapMobileNo := make(map[string]int)
	for _,input := range inputs{
		input = strings.ReplaceAll(input," ","")
		input = strings.ReplaceAll(input,"(","")
		input = strings.ReplaceAll(input,")","")
		input = strings.ReplaceAll(input,"-","")
		val := mapMobileNo[input]
		val++
		mapMobileNo[input] = val;
	}
	return mapMobileNo
}