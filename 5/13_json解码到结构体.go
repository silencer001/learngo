package main

import (
	"encoding/json"
	"fmt"
)

//成员变量名首字母必须大写
type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOK     bool     //`json:",string"`
	Price    float64  //`json:",string"` //-表示字段不会输出
}

func main() {
	jsonBuf := `
	{
        "company": "iscaset",
        "isOK": true,
        "price": 66.66,
        "subjects": [
                "Go",
                "C++",
                "Python"
        ]
	}`
	var tmp IT
	err := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		fmt.Println("error string: ", err)
		return
	}
	fmt.Printf("tmp = %+v", tmp)
}
