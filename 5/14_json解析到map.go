package main

import (
	"encoding/json"
	"fmt"
)

//成员变量名首字母必须大写
// type IT struct {
// 	Company  string   `json:"company"`
// 	Subjects []string `json:"subjects"`
// 	IsOK     bool     //`json:",string"`
// 	Price    float64  //`json:",string"` //-表示字段不会输出
// }

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
	//创建一个map
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		fmt.Println("error string: ", err)
		return
	}
	fmt.Printf("tmp = %+v\n", m)

	// cannot use m["company"] (type interface {}) as type string in assignment: need type assertion
	//str = m["company"]
	//必须要使用类型断言
	for k, v := range m {
		fmt.Printf("%v ========== %v, type %T\n", k, v, v)
		if str, ok := v.(string); ok {
			fmt.Printf("str == %s\n", str)
		} else if slice, ok := v.([]interface{}); ok {
			fmt.Printf("slice == %v\n", slice)
		}
	}
}
