package util

import (
	"fmt"
	"strconv"
)

func Query(username string, password string) []*Course {
	query := newQuery(GetReady(), username, password)
	return GetData(query)
}
func NewCourses(course []*Course) *Courses {
	var ret Courses
	ret.cre = course
	ret.total = len(course)
	ret.valid = ret.total
	ret.average = 0.0
	ret.gpa = 0.0
	return &ret
}
func Handle(course []*Course, key []string) *Courses {
	m := NewCourses(course)
	sumScore := 0.0
	sumCredit := 0.0
	sumPoint := 0.0
	valid := 0
	for _, v := range course {
		credit, _ := strconv.ParseFloat(string(v.credit), 64)
		score, _ := strconv.ParseFloat(string(v.score), 64)
		if v.category == "001" && !isContain(v.name, key) {
			valid++
			sumCredit += credit
			sumScore += score * credit
		}
		sumPoint += v.point * credit
	}
	m.valid = valid
	m.average = sumScore / sumCredit
	m.gpa = sumPoint / sumCredit
	return m
}
func (self *Courses) Info() {
	for _, v := range self.cre {
		fmt.Println(*v)
	}
	fmt.Printf("average: %.3f\n", self.average)
	fmt.Printf("gpa: %.3f\n", self.gpa)
}
