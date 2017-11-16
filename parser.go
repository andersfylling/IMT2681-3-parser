package parser

import (
	"errors"
	"strings"
	"log"
	"encoding/json"
	"io/ioutil"
	//		"fmt"
)

// ExchangeRate holds teh base and target currencies for rate requests
type ExchangeRate struct {
	Base   string
	Target string
}

// ParseStr takes a str arg and returns nil error if a base and target currency was found
// the base and target is then saved to the struct returned.
// See tests for logic.
func ParseStr(input string) (*ExchangeRate, error) {
	res := &ExchangeRate{
		Base:   "",
		Target: "",
	}

	// Put parser logic here
	input = strings.ToLower(input)
	match := strings.Split(input, "and")

	//fmt.Println(match)

	sheetData, err := ioutil.ReadFile("abbreviations.json")
	if err != nil {
		log.Fatalln(err)
	}
	var m map[string]string
	json.Unmarshal(sheetData, &m)

	if len(match) == 2 {
		baseSat := false
		tarSat := false
		for a, b := range m {
			//fmt.Printf("%v: %v", a, b)
			//fmt.Println()
			pattern := []string{a}
			if strings.Contains(match[0], b) {
				res.Base = strings.ToUpper(a)
				baseSat = true
			}

			if strings.Contains(match[1], b)  {
				res.Target = strings.ToUpper(a)
				tarSat = true
			}
			baseMatch := Search(match[0], pattern)
			tarMatch := Search(match[1], pattern)
			if len(baseMatch) > 0 && baseSat == false {
				res.Base = strings.ToUpper(a)
				baseSat = true
			}
			if len(tarMatch) > 0 && tarSat == false{
				res.Target = strings.ToUpper(a)
				tarSat = true
			}
		}
	}

	// check for missing data
	if res.Base == "" || res.Target == "" {
		return res, errors.New("Could not identify it as a currency rate question")
	}

	return res, nil
}
