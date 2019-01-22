package main

import (
	"flag"
	"fmt"
	"util"
)

var (
	help     bool
	username string
	password string
	ignore   util.IgnoreString
	confirm  bool
)

func main() {
	flag.Parse()
	if help || username == "" || password == "" {
		flag.Usage()
		return
	}
	if !confirm {
		fmt.Println("you should use `-y` to pass the validation after you have checked the `-u, -p`, " +
			"because the username and password must be right")
		return
	}
	data := util.Query(username, password)
	util.Handle(data, []string(ignore)).Info()
}

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.StringVar(&username, "u", "", "username")
	flag.StringVar(&password, "p", "", "password")
	flag.Var(&ignore, "i", "ignore key words(use `,` to split)")
	flag.BoolVar(&confirm, "y", false, "confirm the username and password")
}
