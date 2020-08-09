package jsondef

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo `json:"basic_info"`
	JobInfo   `json:"job_info"`
}
