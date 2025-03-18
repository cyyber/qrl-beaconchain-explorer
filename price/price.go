package price

var currencies = map[string]struct {
	Symbol string
	Label  string
}{
	"ZND": {"ZND", "Zond ZND"},
}

func GetPrice(a, b string) float64 {
	return 1.0
}

func GetCurrencySymbol(currency string) string {
	x, exists := currencies[currency]
	if !exists {
		return ""
	}
	return x.Symbol
}
