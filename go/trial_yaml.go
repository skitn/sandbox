package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Member struct {
	Members []string
}

func main() {
	buf, err := ioutil.ReadFile("trial_yaml.yml")
	if err != nil {
		panic(err)
	}

	var member Member
	err = yaml.Unmarshal(buf, &member)
	fmt.Println(member)
}
