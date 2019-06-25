package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string `json:vip_name`
	Age int `json:"age"`
}

func main(){
	str:=[]byte(`{"name":"vip_whisper","age":22,"ok":1}`)
	a:=A{}

	err:=json.Unmarshal(str,&a)
	//fmt.Println("a=",a)
	//rs,err:=json.Marshal(a)
	if err!=nil {
		fmt.Println("failed")
	}else {
		fmt.Println("a=",a)
		//fmt.Println("string=",string(a))
	}

	b:=A{"a",13}
	index,_:=json.Marshal(b)
	fmt.Println(string(index))
}


