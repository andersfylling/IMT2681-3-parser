package parser

func hasNrOfCurrencies(c *ExchangeRate) int {
  total := 0
  if len(c.Base) > 0 {
    total += 1
  }
  if len(c.Target) > 0 {
    total += 1
  }
  
  return total
}

func reassureNrOfCurrencies(inputs []string, expected int) {
  logrus.Info("Testing with %d-currency sentences", expected)
  for _, input := range inputs {
    res, err := ParseStr(input)
    got := hasNrOfCurrencies(res)
    if expected < 3 {
      if err != nil {
        t.Errorf("Got an error when expecting none. Error: "err.Error())
      } else if got != expected {
        t.Errorf("Incorrect number of detected currencies, got: %d, want: %d.", got, expected)
      }
    } else if err == nil {
        t.Errorf("Did not get an error when there were more than three mentioned currencies. Error: "err.Error())
    }
  }
}

func TestParseStr(*testing.T) {
  // Valid sentences with only two currencies
  valids := []string{
    "What is the current exchange rate between Norwegian Kroner and Euro?",
    "What is the exchange rate between USD and NOK?",
    "What is the exchange rate between euro and norwegian kroner?",
  }
  
  // Sentences which contains no currencies
  sentencesNone := []string{
    "What is a dinosaur?",
  }
  
  // Sentences with one currency
  sentencesOne := []string{
    "Lol NOK",
  }
  
  // Sentences with more than two currencies
  sentencesThree := []string{
    "NOK, NOK, NOK",
  }


  logrus.Info("Testing with valid sentences")
  reassureNrOfCurrencies(valids, 2)
  
  logrus.Info("Testing with no-currency sentences")
  reassureNrOfCurrencies(sentencesNone, 0)
  
  logrus.Info("Testing with one-currency sentences")
  reassureNrOfCurrencies(sentencesOne, 1)
  
  logrus.Info("Testing with one-currency sentences")
  reassureNrOfCurrencies(sentencesThree, 3)
  
  logrus.Info("Done")
}
