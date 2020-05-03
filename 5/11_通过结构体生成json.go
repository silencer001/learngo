package main

import (
	"encoding/json"
	"fmt"
)

/*
{
    "Company":"itcaset",
    "Subjects":[
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "IsOK":true,
    "Price":66.66
}
*/
//成员变量名首字母必须大写
type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOK     bool     `json:",string"`
	Price    float64  `json:"-"` //-表示字段不会输出
}

func main() {
	s := IT{"itcaset", []string{"Go", "C++", "Python", "Test"}, true, 66.66}

	//编码，根据内容生成json文本
	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "#", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println("buf = ", string(buf))
}
