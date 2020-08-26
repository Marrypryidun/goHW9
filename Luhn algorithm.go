package main

import (
	"fmt"
)

func algorithmLuhn(s string)  bool{
	var num int = 0;
	var sum int=0;
	var second = len(s) % 2
	for i:=len(s)-1 ;i>=0;i-- {
		num = int(s[i]-'0')
		if(i%2==second) {
			num *= 2
			//fmt.Println("ghsd",s[i]-'0')
			if (num > 9) {
				num = num%10+1
			}
		}
		//fmt.Println(num)
		sum += num
	}
	fmt.Println(sum)
	return sum%10 == 0
}
func answer(b bool)  {
	if(b){
		println("The number is correct")
	}else{
		println("The number is incorrect")
	}
}
func main()  {
	answer(algorithmLuhn("79927398713"))
	answer(algorithmLuhn("79927398710"))
}