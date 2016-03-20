package main

import (
	"errors"
	"fmt"
)

// AwesomeEngineer represents an awesome engineer.
type AwesomeEngineer struct {
	company string
}

func (e *AwesomeEngineer) join(newCompany string) {
	e.company = newCompany
}

func (e *AwesomeEngineer) impact() (string, error) {
	if e.company == "Thumbtack" {
		return "Help people accomplish projects.", nil
	}
	return "", errors.New("Not enough impact. Email ni...@thumbtack.com to fix this error")
}

func main() {
	you := AwesomeEngineer{}
	you.join("Thumbtack")
	impact, err := you.impact()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(impact)
}
