package einvalidator

import (
	"testing"
)

func TestValidate(t *testing.T) {
	einV := New()
	cases := []struct {
		ein   string
		valid bool
	}{
		{"20-0000000", true},
		{"20 0000000", true},
		{"200000000", true},
		{"20x0000000", false}, // invalid seperator
		{"97-9999999", false}, // invalid prefix
		{"", false}, // no ein passed
		{"00", false}, // only prefix passed
		{"not a valid EIN", false}, // not even close to EIN passed
	}
	for _, c := range cases {
		valid, err := einV.Validate(c.ein)
		if c.valid != valid {
			t.Errorf("failed to validate: %s should be %t, but received error %s", c.ein, c.valid, err)
		}
	}
}

func TestMask(t *testing.T) {
	einV := New()
	cases := []struct {
		ein   string
		mask  string
	}{
		{"20-0000000", "XX-XXXXX00"},
		{"20 0000000", "XX XXXXX00"},
		{"200000000", "XXXXXXX00"},
		{"20y0000000", ""}, // invalid seperator
		{"97-9999999", ""}, // invalid prefix
		{"", ""}, // no ein passed
		{"00", ""}, // only prefix passed
		{"not a valid EIN", ""}, // not even close to EIN passed
	}
	for _, c := range cases {
		mask, err := einV.Mask(c.ein)
		if c.mask != mask {
			t.Errorf("failed to validate: %s should be %s, but received error %s", c.ein, c.mask, err)
		}
	}
}
