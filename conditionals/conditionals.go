package conditionals

import (
	// "errors"
	"fmt"
	"strconv"
)

func myFunc(val int) (result int, err error){
	// err = errors.New("this is an error")
	err = nil
	return val, err
}

func Conditionals(){
	a := 34
	b := 45
	if a<b{
		fmt.Println(a<b)
	}else{
		fmt.Println("Not True")
	}

	var result, err_ = myFunc(2)

	fmt.Println(err_.Error(), result)


	n := 345

	if n%2 == 0{
		fmt.Println("N is even")
	}else{
		fmt.Println("N is odd")
	}

	stringval := "5"

	var value, err = strconv.Atoi(stringval)

	if(err != nil){
		println("This isn't an integrer")
	}else{
		println(value)
	}
}