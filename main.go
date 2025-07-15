package main

import (
	"fmt"
)

func main(){
	input:="d3:cow3:moo4:spam4:eggse"
	result,err=decoder.Decode([]byte(input))
	if(err!=nil){
		fmt.Println("Error in Decoding : %v\n",err)
		return
	}
	fmt.Printf("Decoded: %#v\n",result)
	encoded,err:=encoder.Encode(result)
	if(err!=nil){
		fmt.Println("Error in Encoding: %v\n",err)
		return
	}
	fmt.Printf("Encoded: %s\n"string(encoded))

}
