package utils

import "fmt"

func DisplayError(err error){
	if err != nil{
		fmt.Println(err)
	}
}