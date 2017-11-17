package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/yut-kt/goholiday"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Custom_holidays []string
}

func main() {
	buf, err := ioutil.ReadFile("trial_holiday.yml")
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

	holidays := []time.Time{}
	for i := 0; i < len(config.Custom_holidays); i++ {
		datetime, err := time.Parse("2006-01-02", config.Custom_holidays[i])
		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
			return
		}
		holidays = append(holidays, datetime)
	}

	goholiday.SetUniqueHolidays(holidays)
	fmt.Println(holidays)

	// Check customize holiday
	datetime2, err := time.Parse("2006-01-02", "2017-11-17")
	fmt.Println(goholiday.IsNationalHoliday(datetime2))
}
