package parser

// ExchangeRate holds teh base and target currencies for rate requests
type ExhangeRate struct {
  Base string
  Target string
}

// ParseStr takes a str arg and returns nil error if a base and target currency was found
// the abse and target is then saved to the struct returned.
func ParseStr(input string) (*ExchangeRate, error) {
  res := &ExchangeRate{
    Base: "", 
    Target: "",
  }
  
  // Put parser logic here
  
  // check for missing data
  if res.Base == "" || res.Target == "" {
    return res, errors.New("Could not identify it as a currency rate question")
  }
  
  return res, nil
}
