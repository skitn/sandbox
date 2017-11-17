package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Members []string
}

func main() {
	buf, err := ioutil.ReadFile("trial_yaml.yml")
	if err != nil {
		fmt.Println(fmt.Errorf("error: %s", err))
		return
	}

	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %s", err))
		return
	}

	for i := 0; i < len(config.Members); i++ {
		fmt.Println(config.Members[i])
	}
}
