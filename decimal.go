package decimal

import (
	"encoding/json"
	"fmt"
	"math"
)

const (
	DEFAULT_DECIMAL_POINT      = ','
	DEFAULT_THOUSAND_SEPARATOR = '.'
)

type Decimal struct {
	whole             int64
	fraction          int64
	precision         uint
	decimalPoint      byte
	thousandSeparator byte
}

// NewDecimal - Creates a new decimal from an integer
// Amount must be an integer representing the decimal as decimal * 10^-1*precision.
// For example, if the decimal is 123.45, amount must be 12345 and precision 2
func NewDecimal(amount int64, precision uint) *Decimal {
	// Convert amount to float by multiplying it with 10 to the power of negative precision
	floatAmount := float64(amount) * math.Pow10(int(precision)*-1)
	// Split the float into its whole and fractional part
	w, f := math.Modf(floatAmount)
	// Determine the sign of the number. We need this for rounding the fractional part below
	sign := 1.0
	if w < 0 || f < 0 {
		sign = -1.0
	}
	// Rounder is used to round the fraction part. We multiply 5 by sign by 10 to the power of
	// negative precision incremented by one. If the fraction is 0.456 and the precision is 2
	// the fraction will be converted to an integer as 46
	rounder := sign * (5 * math.Pow10(int(precision+1)*-1))
	return &Decimal{
		whole:             int64(w),
		fraction:          int64((f + rounder) * math.Pow10(int(precision))),
		precision:         precision,
		decimalPoint:      DEFAULT_DECIMAL_POINT,
		thousandSeparator: DEFAULT_THOUSAND_SEPARATOR,
	}
}

// NewDecimalFromFloat - Creates a new decimal from a float
// Amount must be a float and precision the desired width of the fraction part.
// For example, amount = 123.45 and precision 2
func NewDecimalFromFloat(amount float64, precision uint) *Decimal {
	// Split the float into its whole and fractional part
	w, f := math.Modf(amount)
	nf := f * math.Pow10(int(precision+2))
	rnf := math.Round(nf)
	f = rnf * math.Pow10(int(precision+2)*-1)
	// Determine the sign of the number. We need this for rounding the fractional part below
	sign := 1.0
	if w < 0 || f < 0 {
		sign = -1.0
	}
	// Rounder is used to round the fraction part. We multiply 5 by sign by 10 to the power of
	// negative precision incremented by one. If the fraction is 0.456 and the precision is 2
	// the fraction will be converted to an integer as 46
	rounder := sign * (5 * math.Pow10(int(precision+1)*-1))
	return &Decimal{
		whole:             int64(w),
		fraction:          int64((f + rounder) * math.Pow10(int(precision))),
		precision:         uint(precision),
		decimalPoint:      DEFAULT_DECIMAL_POINT,
		thousandSeparator: DEFAULT_THOUSAND_SEPARATOR,
	}
}

// ToInt - Returns the integer representation of decimal multiplied by 10^precision
// For example, for a decimal with whole part = 123 and fraction 45, the return
// value will be 12345
func (d Decimal) ToInt() int64 {
	w := d.whole * int64(math.Pow10(int(d.precision)))
	return w + d.fraction
}

// ToFloat - Returns a float representation of decimal, with number of floating points
// equal to decimal precision. For example, for a decimal with whole part = 123 and fraction 45
// the return value will be 123.45
func (d Decimal) ToFloat() float64 {
	f := float64(d.fraction) * math.Pow10(int(d.precision)*-1)
	return float64(d.whole) + f
}

// ToString - Returns the decimal as a string. For a decimal with whole part = 123 and fraction 45,
// the return value will be '123.45'
func (d Decimal) ToString() string {
	f := d.ToFloat()
	return fmt.Sprintf("%."+fmt.Sprintf("%d", d.precision)+"f", f)
}

// ToStringFormatted - Returns the decimal as a string, formatted with thousands separator and decimal
// point. For a decimal with decimal point '.' and thousands separator ',' and value 2334599 the result
// will be '23,345.99'
func (d Decimal) ToStringFormatted() string {
	// Count the number of times the whole part of the decimal is divisible by 1000
	t := 0
	for r := d.whole / 1000; r > 0; r /= 1000 {
		t++
	}
	// Convert the whole part to string
	fs := fmt.Sprintf("%d", d.whole)
	// Add the thousand seperator every three characters of fs
	for i := 1; i <= t*3; i += 4 {
		fs = fs[:i] + string(d.thousandSeparator) + fs[i:]
	}
	// Add the decimal point
	fs += string(d.decimalPoint)
	// Add the fractional part to the string padding zeroes to the right as required by
	// the decimal precision.
	fs += fmt.Sprintf("%0"+fmt.Sprintf("%d", d.precision)+"d", d.fraction)
	return fs
}

// SetDecimalPoint - Setter for decimal point
func (d *Decimal) SetDecimalPoint(decimalPoint byte) *Decimal {
	d.decimalPoint = decimalPoint
	return d
}

// SetThousandSeparator - Setter for thousand separator
func (d *Decimal) SetThousandSeparator(ts byte) *Decimal {
	d.thousandSeparator = ts
	return d
}

// GetWhole - Getter for the whole part of the decimal
func (d Decimal) GetWhole() int64 {
	return d.whole
}

// GetFraction - Getter for the fraction part of the decimal
func (d Decimal) GetFraction() int64 {
	return d.fraction
}

// IsZero - Returns true if decimal is zero
func (d Decimal) IsZero() bool {
	return d.whole == 0 && d.fraction == 0
}

// IsNegative - Returns true if decimal is negative
func (d Decimal) IsNegative() bool {
	return d.whole < 0 && d.fraction <= 0
}

// Add - Adds a decimal to another decimal. The resulting decimal will have
// the precision of the decimal with the largest precision.
func (d Decimal) Add(decimalToAdd Decimal) Decimal {
	sum := d.ToFloat() + decimalToAdd.ToFloat()
	precision := math.Max(float64(d.precision), float64(decimalToAdd.precision))
	return *NewDecimalFromFloat(sum, uint(precision))
}

// Subtract - Subtracts a decimal from another decimal. The resulting decimal will
// have the precision of the decimal with the largest precision.
func (d Decimal) Subtract(decimalToSubtract Decimal) Decimal {
	difference := d.ToFloat() - decimalToSubtract.ToFloat()
	precision := math.Max(float64(d.precision), float64(decimalToSubtract.precision))
	return *NewDecimalFromFloat(difference, uint(precision))
}

// Multiply - Multiplies a decimal with another decimal. The resulting decimal will
// have the precision of the decimal with the largest precision.
func (d Decimal) Multiply(factor Decimal) Decimal {
	product := d.ToFloat() * factor.ToFloat()
	precision := math.Max(float64(d.precision), float64(factor.precision))
	return *NewDecimalFromFloat(product, uint(precision))
}

// Divide - Divides a decimal with another decimal. The resulting decimal will
// have the precision of the decimal with the largest precision.
func (d Decimal) Divide(divisor Decimal) Decimal {
	quotient := d.ToFloat() / divisor.ToFloat()
	precision := math.Max(float64(d.precision), float64(divisor.precision))
	return *NewDecimalFromFloat(quotient, uint(precision))
}

// AddInt - Adds an integer to decimal
func (d Decimal) AddInt(intToAdd int64) Decimal {
	intToDec := NewDecimal(intToAdd, 0)
	return d.Add(*intToDec)
}

// SubtractInt - Subtracts and integer from a decimal
func (d Decimal) SubtractInt(intToSubtract int64) Decimal {
	intToDec := NewDecimal(intToSubtract, 0)
	return d.Subtract(*intToDec)
}

// MultiplyByInt - Multiplies a decimal by an integer
func (d Decimal) MultiplyByInt(intToMultiply int64) Decimal {
	intToDec := NewDecimal(intToMultiply, 0)
	return d.Multiply(*intToDec)
}

// DivideByInt - Divides a decimal by an integer
func (d Decimal) DivideByInt(inToDivide int64) Decimal {
	intToDec := NewDecimal(inToDivide, 0)
	return d.Divide(*intToDec)
}

func (d Decimal) AddFloat(floatToAdd float64) Decimal {
	floatToDec := NewDecimalFromFloat(floatToAdd, d.precision)
	return d.Add(*floatToDec)
}

func (d Decimal) SubtractFloat(floatToSubtract float64) Decimal {
	floatToDec := NewDecimalFromFloat(floatToSubtract, d.precision)
	return d.Subtract(*floatToDec)
}

func (d Decimal) MultiplyFloat(floatToMultiply float64) Decimal {
	decimalToFloat := d.ToFloat()
	product := decimalToFloat * floatToMultiply
	return *NewDecimalFromFloat(product, d.precision)
}

func (d Decimal) DivideByFloat(floatToDivideBy float64) Decimal {
	decimalToFloat := d.ToFloat()
	quotient := decimalToFloat / floatToDivideBy
	return *NewDecimalFromFloat(quotient, d.precision)
}

func (d Decimal) Split(toParts uint) []Decimal {
	decimalToFloat := d.ToFloat()
	parts := make([]Decimal, toParts)
	quotient := decimalToFloat / float64(toParts)
	quotient *= math.Pow10(int(d.precision))
	quotient = math.Trunc(quotient)
	quotient *= math.Pow10(int(d.precision) * -1)
	var sum float64
	for i := uint(0); i < toParts; i++ {
		parts[i] = *NewDecimalFromFloat(quotient, d.precision)
		sum += quotient
	}
	remainder := decimalToFloat - sum
	parts[0] = parts[0].AddFloat(remainder)
	return parts
}

// MarshalJSON -
func (d *Decimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ToFloat())
}

// UnmarshalJSON -
func (d *Decimal) UnmarshalJSON(data []byte) error {
	var floatVal float64
	if err := json.Unmarshal(data, &floatVal); err != nil {
		return fmt.Errorf("unmarshal decimal: %v", err.Error())
	}
	dc := NewDecimalFromFloat(floatVal, d.precision)
	d.whole = dc.whole
	d.fraction = dc.fraction
	return nil
}
