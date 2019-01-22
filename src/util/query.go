package util

import (
	"bytes"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"strconv"
)

func newQuery(id string, username string, password string) *queryFormat {
	return &queryFormat{
		id:       id,
		captcha:  "",
		username: username,
		password: password,
	}
}
func GetReady() string {
	req, _ := http.Get(getReady)
	data, _ := ioutil.ReadAll(req.Body)
	ret, _ := simplejson.NewJson(data)
	id := ret.Get("id").MustFloat64()
	return strconv.FormatFloat(id, 'f', 0, 64)
}

func GetData(format *queryFormat) []*Course {
	var courses []*Course
	postData := query2byte(format)
	req, _ := http.NewRequest("POST", getData, bytes.NewReader(postData))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	all, _ := simplejson.NewJson(body)
	res, err := all.Get("result").String()
	if err != nil {
		return nil
	}
	if res == "login_success" {
		data, _ := all.Get("data").Get("rows").Array()
		for _, v := range data {
			course := v.(map[string]interface{})
			m := con2course(course)
			courses = append(courses, m)
		}

	} else {
		fmt.Println("error: ", res)
		return nil
	}
	return courses
}
