package util

import "encoding/json"

const (
	getReady string = "https://api.ee.xidian.edu.cn/ehall/get_ready.json"
	getData  string = "https://api.ee.xidian.edu.cn/ehall/get_data.json"
)

type queryFormat struct {
	id       string
	captcha  string
	username string
	password string
}

type Course struct {
	name     string
	score    json.Number
	credit   json.Number
	category string
	point    float64
}
type Courses struct {
	cre     []*Course
	total   int
	valid   int
	gpa     float64
	average float64
}

type IgnoreString []string
