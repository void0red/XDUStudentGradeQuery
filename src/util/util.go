package util

import (
	"encoding/json"
	"strconv"
	"strings"
)

func query2byte(format *queryFormat) []byte {
	m := make(map[string]interface{})
	m["id"] = format.id
	m["captcha"] = format.captcha
	m["username"] = format.username
	m["password"] = format.password
	data, _ := json.Marshal(&m)
	return data
}
func con2course(data map[string]interface{}) *Course {
	ret := &Course{
		name:     data["KCM"].(string),
		score:    data["ZCJ"].(json.Number),
		credit:   data["XF"].(json.Number),
		category: data["KCXZDM_DISPLAY"].(string),
	}
	score, _ := strconv.ParseFloat(string(data["ZCJ"].(json.Number)), 64)
	ret.point = score2gpa(score)
	return ret
}
func isContain(target string, key []string) bool {
	for _, v := range key {
		if strings.Contains(target, v) {
			return true
		}
	}
	return false
}
func score2gpa(score float64) float64 {
	if score >= 95 {
		return 4.0
	} else if score >= 90 {
		return 3.9
	} else if score >= 84 {
		return 3.8
	} else if score >= 80 {
		return 3.6
	} else if score >= 76 {
		return 3.4
	} else if score >= 73 {
		return 3.2
	} else if score >= 70 {
		return 3.0
	} else if score >= 67 {
		return 2.7
	} else if score >= 64 {
		return 2.4
	} else if score >= 62 {
		return 2.2
	} else if score >= 60 {
		return 2.0
	} else {
		return 0.0
	}
}

func (self *IgnoreString) Set(val string) error {
	*self = strings.Split(val, ",")
	return nil
}
func (self *IgnoreString) String() string {
	return strings.Join([]string(*self), ",")
}
