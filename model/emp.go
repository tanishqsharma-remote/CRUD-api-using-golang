package model

type Emp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type EmpList struct {
	Items []Emp `json:"emps"`
}
