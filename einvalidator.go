package einvalidator // import "src.techknowlogick.com/einvalidator"

import (
	"fmt"
	"regexp"
)

var (
	DefaultEINValidator = New()
	einRegex            = regexp.MustCompile(`^\d{2}[- ]{0,1}\d{7}$`)
	replaceRegex        = regexp.MustCompile(`[\w]`)
	// prefixes found here: https://www.irs.gov/businesses/small-businesses-self-employed/how-eins-are-assigned-and-valid-ein-prefixes
	campuses = map[string][]string{ // prefixes stored as strings as they can be 0 prefixed
		"andover":      []string{"10", "12"},
		"atlanta":      []string{"60", "67"},
		"austin":       []string{"50", "53"},
		"brookhaven":   []string{"01", "02", "03", "04", "05", "06", "11", "13", "14", "16", "21", "22", "23", "25", "34", "51", "52", "54", "55", "56", "57", "58", "59", "65"},
		"cincinnati":   []string{"30", "32", "35", "36", "37", "38", "61"},
		"fresno":       []string{"15", "24"},
		"kansas_city":  []string{"40", "44"},
		"memphis":      []string{"94", "95"},
		"ogden":        []string{"80", "90"},
		"philadelphia": []string{"33", "39", "41", "42", "43", "46", "48", "62", "63", "64", "66", "68", "71", "72", "73", "74", "75", "76", "77", "82", "83", "84", "85", "86", "87", "88", "91", "92", "93", "98", "99"},
		"internet":     []string{"20", "26", "27", "45", "46", "47", "81", "82", "83"},
		"sba":          []string{"31"}, // Small Business Administration
	}
)

type EINValidator struct {
}

func New() *EINValidator {
	return &EINValidator{}
}

func (v *EINValidator) Validate(ein string) (bool, error) {
	// Two failiures: 1, invalid format 2, invalid campus
	if !einRegex.Match([]byte(ein)) {
		return false, fmt.Errorf("invalid format of provided EIN")
	}
	for _, prefixes := range campuses {
		for _, prefix := range prefixes {
			if ein[0:2] == prefix { // matches first two chars of ein to prefix
				return true, nil
			}
		}
	}
	return false, fmt.Errorf("not using valid prefix")
}

func Validate(ein string) (bool, error) {
	return DefaultEINValidator.Validate(ein)
}

func (v *EINValidator) Mask(ein string) (string, error) {
	if ok, err := v.Validate(ein); !ok {
		return "", err // return blank string to prevent any data leakage
	}
	firstPart := ein[:len(ein)-2]
	lastTwo := ein[len(ein)-2:]
	return replaceRegex.ReplaceAllString(firstPart, "X") + lastTwo, nil
}

func Mask(ein string) (string, error) {
	return DefaultEINValidator.Mask(ein)
}
