package decimal_test

import (
	"testing"

	"github.com/petrossordinas/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewDecimal(t *testing.T) {
	tests := []struct {
		intVal    int64
		precision uint
		whole     int64
		fraction  int64
	}{
		{2456, 2, 24, 56},
		{12357, 3, 12, 357},
		{314159265358979, 14, 3, 14159265358979},
		{38948737383, 4, 3894873, 7383},
		{18, 2, 0, 18},
		{271828, 5, 2, 71828},
		{0, 2, 0, 0},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimal(tc.intVal, tc.precision)
		assert.EqualValues(tc.whole, d.GetWhole(), "Test No: %d - Should be equal", testNo+1)
		assert.EqualValues(tc.fraction, d.GetFraction(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestNewDecimalFromFloat(t *testing.T) {
	tests := []struct {
		floatVal  float64
		precision uint
		whole     int64
		fraction  int64
	}{
		{24.56, 2, 24, 56},
		{12.357, 3, 12, 357},
		{3.14159265358979, 14, 3, 14159265358979},
		{3894873.7383, 4, 3894873, 7383},
		{0.18, 2, 0, 18},
		{2.71828, 5, 2, 71828},
		{0.0, 2, 0, 0},
		{2.675, 2, 2, 68},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(tc.floatVal, tc.precision)
		assert.EqualValues(tc.whole, d.GetWhole(), "Test No: %d - Should be equal", testNo+1)
		assert.EqualValues(tc.fraction, d.GetFraction(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestDecimalToInt(t *testing.T) {
	tests := []struct {
		floatVal  float64
		precision uint
		intVal    int64
	}{
		{24.56, 2, 2456},
		{12.357, 3, 12357},
		{3.14159265358979, 14, 314159265358979},
		{3894873.7383, 4, 38948737383},
		{0.18, 2, 18},
		{2.71828, 5, 271828},
		{0.0, 2, 0},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(tc.floatVal, tc.precision)
		assert.EqualValues(tc.intVal, d.ToInt(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestDecimalToFloat(t *testing.T) {
	tests := []struct {
		intVal    int64
		precision uint
		floatVal  float64
	}{
		{2456, 2, 24.56},
		{12357, 3, 12.357},
		{314159265358979, 14, 3.14159265358979},
		{38948737383, 4, 3894873.7383},
		{18, 2, 0.18},
		{271828, 5, 2.71828},
		{0, 2, 0.0},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimal(tc.intVal, tc.precision)
		assert.EqualValues(d.ToFloat(), tc.floatVal, "Test No: %d - Should be equal", testNo+1)
	}
}

func TestDecimalToString(t *testing.T) {
	tests := []struct {
		intVal    int64
		precision uint
		strVal    string
	}{
		{2456, 2, "24.56"},
		{12357, 3, "12.357"},
		{314159265358979, 14, "3.14159265358979"},
		{38948737383, 4, "3894873.7383"},
		{18, 2, "0.18"},
		{271828, 5, "2.71828"},
		{0, 2, "0.00"},
		{0, 1, "0.0"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimal(tc.intVal, tc.precision)
		assert.EqualValues(tc.strVal, d.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestDecimalToStringFormatted(t *testing.T) {
	tests := []struct {
		intVal    int64
		precision uint
		strVal    string
	}{
		{2456, 2, "24,56"},
		{12357, 3, "12,357"},
		{314159265358979, 14, "3,14159265358979"},
		{38948737383, 4, "3.894.873,7383"},
		{18, 2, "0,18"},
		{271828, 5, "2,71828"},
		{0, 2, "0,00"},
		{0, 1, "0,0"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimal(tc.intVal, tc.precision)
		assert.EqualValues(tc.strVal, d.ToStringFormatted(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestDecimalToStringFormattedWithChangingSeparators(t *testing.T) {
	tests := []struct {
		intVal            int64
		precision         uint
		thousandSeparator byte
		decimalPoint      byte
		strVal            string
	}{
		{2456, 2, '.', ',', "24,56"},
		{12357, 3, ',', '.', "12.357"},
		{314159265358979, 14, '.', ',', "3,14159265358979"},
		{38948737383, 4, ',', '.', "3,894,873.7383"},
		{18, 2, '.', ',', "0,18"},
		{271828, 5, '.', ',', "2,71828"},
		{0, 2, ',', '.', "0.00"},
		{0, 1, '.', ',', "0,0"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimal(tc.intVal, tc.precision)
		d.SetDecimalPoint(tc.decimalPoint)
		d.SetThousandSeparator(tc.thousandSeparator)
		assert.EqualValues(tc.strVal, d.ToStringFormatted(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestIsZero(t *testing.T) {
	tests := []struct {
		intVal    int64
		precision uint
		result    bool
	}{
		{2456, 2, false},
		{12357, 3, false},
		{0, 2, true},
		{0, 1, true},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimal(tc.intVal, tc.precision)
		assert.Equal(tc.result, d.IsZero(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestIsNegative(t *testing.T) {
	tests := []struct {
		intVal    int64
		precision uint
		result    bool
	}{
		{2456, 2, false},
		{-12357, 3, true},
		{0, 2, false},
		{-4990, 1, true},
		{0, 0, false},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimal(tc.intVal, tc.precision)
		assert.Equal(tc.result, d.IsNegative(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		decimal1   float64
		precision1 uint
		decimal2   float64
		precision2 uint
		result     string
	}{
		{24.56, 2, 5.44, 2, "30.00"},
		{12.34, 2, 35.0, 4, "47.3400"},
		{15.5, 2, 17.958, 3, "33.458"},
		{10.75, 2, -15.75, 2, "-5.00"},
		{-10.7, 2, -4.3, 5, "-15.00000"},
		{0.0, 3, 17.5, 1, "17.500"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d1 := decimal.NewDecimalFromFloat(tc.decimal1, tc.precision1)
		d2 := decimal.NewDecimalFromFloat(tc.decimal2, tc.precision2)
		sum := d1.Add(*d2)
		assert.Equal(tc.result, sum.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		decimal1   float64
		precision1 uint
		decimal2   float64
		precision2 uint
		result     string
	}{
		{24.56, 2, 5.44, 2, "19.12"},
		{12.34, 2, 35.0, 4, "-22.6600"},
		{15.5, 2, 17.958, 3, "-2.458"},
		{10.75, 2, -15.75, 2, "26.50"},
		{-10.7, 2, -4.3, 5, "-6.40000"},
		{0.0, 3, 17.5, 1, "-17.500"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d1 := decimal.NewDecimalFromFloat(tc.decimal1, tc.precision1)
		d2 := decimal.NewDecimalFromFloat(tc.decimal2, tc.precision2)
		sum := d1.Subtract(*d2)
		assert.Equal(tc.result, sum.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		decimal1   float64
		precision1 uint
		decimal2   float64
		precision2 uint
		result     string
	}{
		{24.56, 2, 5.44, 2, "133.61"},
		{12.34, 2, 35.0, 4, "431.9000"},
		{15.5, 2, 17.958, 3, "278.349"},
		{10.75, 2, -15.75, 2, "-169.31"},
		{-10.7, 2, -4.3, 5, "46.01000"},
		{0.0, 3, 17.5, 1, "0.000"},
		{0.81, 2, 1.23, 2, "1.00"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d1 := decimal.NewDecimalFromFloat(tc.decimal1, tc.precision1)
		d2 := decimal.NewDecimalFromFloat(tc.decimal2, tc.precision2)
		sum := d1.Multiply(*d2)
		assert.Equal(tc.result, sum.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		decimal1   float64
		precision1 uint
		decimal2   float64
		precision2 uint
		result     string
	}{
		{24.56, 2, 5.44, 2, "4.51"},
		{12.34, 2, 35.0, 4, "0.3526"},
		{15.5, 2, 17.958, 3, "0.863"},
		{10.75, 2, -15.75, 2, "-0.68"},
		{-10.7, 2, -4.3, 5, "2.48837"},
		{0.0, 3, 17.5, 1, "0.000"},
		{1, 2, 1.23, 2, "0.81"},
		{10, 2, 3, 2, "3.33"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d1 := decimal.NewDecimalFromFloat(tc.decimal1, tc.precision1)
		d2 := decimal.NewDecimalFromFloat(tc.decimal2, tc.precision2)
		sum := d1.Divide(*d2)
		assert.Equal(tc.result, sum.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestAddInt(t *testing.T) {
	tests := []struct {
		decimal   float64
		precision uint
		intToAdd  int64
		result    string
	}{
		{24.56, 2, 5, "29.56"},
		{12.34, 2, 35, "47.34"},
		{15.5, 2, 17, "32.50"},
		{10.75, 2, -15, "-4.25"},
		{-10.7, 2, -4, "-14.70"},
		{0.0, 3, 17, "17.000"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(tc.decimal, tc.precision)
		sum := d.AddInt(tc.intToAdd)
		assert.Equal(tc.result, sum.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestSubtractInt(t *testing.T) {
	tests := []struct {
		decimal       float64
		precision     uint
		intToSubtract int64
		result        string
	}{
		{24.56, 2, 5, "19.56"},
		{12.34, 2, 35, "-22.66"},
		{15.5, 2, 17, "-1.50"},
		{10.75, 2, -15, "25.75"},
		{-10.7, 2, -4, "-6.70"},
		{0.0, 3, 17, "-17.000"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(tc.decimal, tc.precision)
		sum := d.SubtractInt(tc.intToSubtract)
		assert.Equal(tc.result, sum.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestMultiplyByInt(t *testing.T) {
	tests := []struct {
		decimal       float64
		precision     uint
		intToMultiply int64
		result        string
	}{
		{24.56, 2, 5, "122.80"},
		{12.34, 2, 35, "431.90"},
		{15.5, 2, 17, "263.50"},
		{10.75, 2, -15, "-161.25"},
		{-10.7, 2, -4, "42.80"},
		{0.0, 3, 17, "0.000"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(tc.decimal, tc.precision)
		sum := d.MultiplyByInt(tc.intToMultiply)
		assert.Equal(tc.result, sum.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestDivideByInt(t *testing.T) {
	tests := []struct {
		decimal     float64
		precision   uint
		intToDivide int64
		result      string
	}{
		{24.56, 2, 5, "4.91"},
		{12.34, 2, 35, "0.35"},
		{15.5, 2, 17, "0.91"},
		{10.75, 2, -15, "-0.72"},
		{10.7, 2, 4, "2.68"},
		{-10.7, 2, -4, "2.68"},
		{0.0, 3, 17, "0.000"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(tc.decimal, tc.precision)
		sum := d.DivideByInt(tc.intToDivide)
		assert.Equal(tc.result, sum.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestAddFloat(t *testing.T) {
	tests := []struct {
		decimal    float64
		precision  uint
		floatToAdd float64
		result     string
	}{
		{24.56, 2, 5.2, "29.76"},
		{12.34, 2, 35.4, "47.74"},
		{15.5, 2, 17.399, "32.90"},
		{10.75, 2, -15.9, "-5.15"},
		{10.7, 2, 4.1, "14.80"},
		{-10.7, 2, -4.1, "-14.80"},
		{0.0, 3, 17.9, "17.900"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(tc.decimal, tc.precision)
		sum := d.AddFloat(tc.floatToAdd)
		assert.Equal(tc.result, sum.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestSubtractFloat(t *testing.T) {
	tests := []struct {
		decimal         float64
		precision       uint
		floatToSubtract float64
		result          string
	}{
		{24.56, 2, 5.2, "19.36"},
		{12.34, 2, 35.4, "-23.06"},
		{15.5, 2, 17.399, "-1.90"},
		{10.75, 2, -15.9, "26.65"},
		{10.7, 2, 4.1, "6.60"},
		{-10.7, 2, -4.1, "-6.60"},
		{0.0, 3, 17.9, "-17.900"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(tc.decimal, tc.precision)
		sum := d.SubtractFloat(tc.floatToSubtract)
		assert.Equal(tc.result, sum.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestMultiplyFloat(t *testing.T) {
	tests := []struct {
		decimal         float64
		precision       uint
		floatToMultiply float64
		result          string
	}{
		{24.56, 2, 5.2, "127.71"},
		{12.34, 2, 35.4, "436.84"},
		{15.5, 2, 17.399, "269.68"},
		{10.75, 2, -15.9, "-170.93"},
		{10.7, 2, 4.1, "43.87"},
		{-10.7, 2, -4.1, "43.87"},
		{0.0, 3, 17.9, "0.000"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(tc.decimal, tc.precision)
		sum := d.MultiplyFloat(tc.floatToMultiply)
		assert.Equal(tc.result, sum.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestDivideByFloat(t *testing.T) {
	tests := []struct {
		decimal         float64
		precision       uint
		floatToDivideBy float64
		result          string
	}{
		{24.56, 2, 5.2, "4.72"},
		{12.34, 2, 35.4, "0.35"},
		{15.5, 2, 17.399, "0.89"},
		{10.75, 2, -15.9, "-0.68"},
		{10.7, 2, 4.1, "2.61"},
		{-10.7, 2, -4.1, "2.61"},
		{0.0, 3, 17.9, "0.000"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(tc.decimal, tc.precision)
		sum := d.DivideByFloat(tc.floatToDivideBy)
		assert.Equal(tc.result, sum.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		decimal    float64
		precision  uint
		splitCount int
		result     []string
	}{
		{10.00, 2, 3, []string{"3.34", "3.33", "3.33"}},
		{14.0, 1, 4, []string{"3.5", "3.5", "3.5", "3.5"}},
		{19.00, 2, 6, []string{"3.20", "3.16", "3.16", "3.16", "3.16", "3.16"}},
		{11.0, 4, 3, []string{"3.6668", "3.6666", "3.6666"}},
		{31.0, 3, 7, []string{"4.432", "4.428", "4.428", "4.428", "4.428", "4.428", "4.428"}},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(tc.decimal, tc.precision)
		dsplit := d.Split(uint(tc.splitCount))
		split := make([]string, tc.splitCount)
		for i, d := range dsplit {
			split[i] = d.ToString()
		}
		assert.Equal(tc.result, split, "Test No: %d - Should be equal", testNo+1)
	}
}

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		decimal   float64
		precision uint
		result    string
	}{
		{35.23, 2, "35.23"},
		{35.32000, 5, "35.32"},
		{148.495049, 6, "148.495049"},
		{23.459, 2, "23.46"},
		{12.009, 2, "12.01"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(tc.decimal, tc.precision)
		json, err := d.MarshalJSON()
		assert.Nil(err, "Was not expecting error")
		assert.Equal(tc.result, string(json), "Test No: %d - Should be equal", testNo+1)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		json      string
		precision uint
		result    string
	}{
		{"35.23", 2, "35.23"},
		{"35.32000", 5, "35.32000"},
		{"148.495049", 2, "148.50"},
		{"23.459", 2, "23.46"},
		{"12.009", 2, "12.01"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(0, tc.precision)
		err := d.UnmarshalJSON([]byte(tc.json))
		assert.Nil(err, "Was not expecting error")
		assert.Equal(tc.result, d.ToString(), "Test No: %d - Should be equal", testNo+1)
	}
}
func TestUnmarshalJSONerrors(t *testing.T) {
	tests := []struct {
		json      string
		precision uint
		result    string
	}{
		{"35.23x", 2, "35.23"},
		{"35 32000", 5, "35.32000"},
	}
	assert := assert.New(t)
	for testNo, tc := range tests {
		d := decimal.NewDecimalFromFloat(0, tc.precision)
		err := d.UnmarshalJSON([]byte(tc.json))
		assert.NotNil(err, "Test No: %d - Was expecting error", testNo+1)
	}
}
