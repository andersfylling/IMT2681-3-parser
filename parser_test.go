package concurrencyparser

import "testing"
import (
	"github.com/sirupsen/logrus"
)

func hasNrOfCurrencies(c *ExchangeRate) int {
	total := 0
	if c.Base != "" {
		total++
	}
	if c.Target != "" {
		total++
	}

	return total
}

func reassureNrOfCurrencies(t *testing.T, inputs []string, expected int) {
	logrus.Infof("Testing with %d-currency sentences", expected)
	for _, input := range inputs {
		res, err := ParseStr(input)
		if err != nil {
			logrus.Info(err.Error())
			return
		}
		got := hasNrOfCurrencies(res)
		if expected == 2 && got == 2 {
			// OK
		} else if expected != 2 {
			t.Errorf("Given currencies were %d, but no error was returned", expected)
		}
	}
}

func reassureBaseAndTargetFields(t *testing.T, inputs []string, expecteds []*ExchangeRate) {
	if len(inputs) != len(expecteds) {
		logrus.Errorf("Length of inputs and expecteds does not match. %d != %d", len(inputs), len(expecteds))
		return
	}

	for i, input := range inputs {
		res, err := ParseStr(input)
		if err != nil {
			logrus.Fatal(err.Error())
			return
		}
		expected := expecteds[i]
		if res.Base != expected.Base || res.Target != expected.Target {
			t.Errorf("Result did not match expected value, got: {%s, %s}, want: {%s, %s}.", res.Base, res.Target, expected.Base, expected.Target)
		}
	}
}

// TestParseStr Should create a exchange rate object per base-target pair found.
func TestParseStr(t *testing.T) {
	// Valid sentences with only two currencies
	valids := []string{
		"What is the current exchange rate between Norwegian Kroner and Euro?",
		"What is the exchange rate between USD and NOK?",
		"What is the exchange rate between euro and norwegian kroner?",
		"NOK and RON (Romanian leu)",
		"What is the exchange rate between danish krone and USD",
		"Exchange rate hkd and myr",
	}
	reassureNrOfCurrencies(t, valids, 2)

	// Sentences which contains no currencies
	sentencesNone := []string{
		"What is a dinosaur? and tull",
	}
	reassureNrOfCurrencies(t, sentencesNone, 0)

	// Sentences with one currency
	sentencesOne := []string{
		"Lol NOK",
	}

	reassureNrOfCurrencies(t, sentencesOne, 1)

	// Sentences with more than two currencies
	sentencesThree := []string{
		"NOK,  NOK, NOK",
	}
	reassureNrOfCurrencies(t, sentencesThree, 3)

	logrus.Info("Done")

	logrus.Info("Now validating correct base and target")
	reassureBaseAndTargetFields(t, valids, []*ExchangeRate{
		&ExchangeRate{Base: "NOK", Target: "EUR"},
		&ExchangeRate{Base: "USD", Target: "NOK"},
		&ExchangeRate{Base: "EUR", Target: "NOK"},
		&ExchangeRate{Base: "NOK", Target: "RON"},
		&ExchangeRate{Base: "DKK", Target: "USD"},
		&ExchangeRate{Base: "HKD", Target: "MYR"},
	})

}
