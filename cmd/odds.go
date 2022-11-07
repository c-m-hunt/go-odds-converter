package cmd

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Odds struct {
	decimalOdds float64
}

type OddsConverter interface {
	ToDecimal() float64
	ToDecimalString() string
	ToFraction() string
	ToUS() float64
	ToUSString() string
	ToImpliedProbability() float64
	ToImpliedProbabilityString() string
	GetReciprocalOdds() Odds
}

func NewOdds(odd string) Odds {
	guess := guessOddsType(odd)
	decimalOdds := 0.0
	switch guess {
	case "decimal":
		decimalOdds, _ = parseDecimal(odd)
	case "fraction":
		decimalOdds, _ = parseFraction(odd)
	case "us":
		decimalOdds, _ = parseUS(odd)
	}
	return Odds{decimalOdds: decimalOdds}
}

func parseDecimal(odd string) (float64, error) {
	return strconv.ParseFloat(odd, 64)
}

func parseFraction(odd string) (float64, error) {
	for _, divider := range []string{SLASH, HYPHEN} {
		if strings.Contains(odd, divider) {
			numbers := strings.Split(odd, divider)
			numerator, _ := strconv.ParseFloat(numbers[0], 64)
			denominator, _ := strconv.ParseFloat(numbers[1], 64)
			return numerator/denominator + 1, nil
		}
	}
	return 0, errors.New("Invalid fraction")
}

func parseUS(odd string) (float64, error) {
	oddConv, err := strconv.ParseFloat(odd, 64)
	if err != nil {
		return 0, err
	}
	if oddConv > 0 {
		return 1 + oddConv/100, nil
	} else if oddConv < 0 {
		return 1 - 100/oddConv, nil
	} else {
		return 0, errors.New("Invalid US odds")
	}
}

func guessOddsType(odd string) string {
	if strings.HasPrefix(odd, "-") || strings.HasPrefix(odd, "+") {
		return US
	}
	for _, divider := range []string{SLASH, HYPHEN} {
		if strings.Contains(odd, divider) {
			return FRACTION
		}
	}
	return DECIMAL
}

func (o Odds) GetReciprocalOdds() Odds {
	return Odds{decimalOdds: (1 / (o.decimalOdds - 1)) + 1}
}

func (o Odds) ToImpliedProbability() float64 {
	return 1 / o.decimalOdds
}

func (o Odds) ToDecimal() float64 {
	return o.decimalOdds
}

func (o Odds) ToDecimalString() string {
	return fmt.Sprintf("%.2f", o.decimalOdds)
}

func (o Odds) ToFraction() string {
	return getFraction(o.decimalOdds-1, SLASH)
}

func (o Odds) ToUS() float64 {
	if o.decimalOdds > 2 {
		return (o.decimalOdds - 1) * 100
	} else {
		return -100 / (o.decimalOdds - 1)
	}
}

func (o Odds) ToUSString() string {
	sign := "+"
	if o.ToUS() < 0 {
		sign = ""
	}
	return fmt.Sprintf("%s%d", sign, int(o.ToUS()))
}

func (o Odds) ToImpliedProbabilityString() string {
	return fmt.Sprintf("%.2f%%", o.ToImpliedProbability()*100)
}

func (o Odds) Display() {
	table := tablewriter.NewWriter(os.Stdout)
	ro := o.GetReciprocalOdds()
	table.Append([]string{"Decimal", o.ToDecimalString()})
	table.Append([]string{"Fraction", o.ToFraction()})
	table.Append([]string{"US", o.ToUSString()})
	table.Append([]string{"Implied Probability", o.ToImpliedProbabilityString()})
	table.Append([]string{"", ""})
	table.Append([]string{"Reciprocal Decimal", ro.ToDecimalString()})
	table.Append([]string{"Reciprocal Fraction", ro.ToFraction()})
	table.Append([]string{"Reciprocal US", ro.ToUSString()})
	table.Append([]string{"Reciprocal Implied Probability", ro.ToImpliedProbabilityString()})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}

func getFraction(dec float64, divider string) string {
	reciprocal := false
	if dec < 1 {
		dec = 1 / dec
		reciprocal = true
	}

	dec = toFixed(dec, 3)
	len := len(strconv.Itoa(int(dec)))
	denom := dec * math.Pow(10, float64(len))
	num := dec * denom
	divisor := gcd(num, denom)
	num = num / divisor
	denom = denom / divisor
	if reciprocal {
		num, denom = denom, num
	}
	return fmt.Sprintf("%d%s%d", int(num), divider, int(denom))
}

func gcd(a float64, b float64) float64 {
	if b < 0.001 {
		return a
	}
	return gcd(b, math.Mod(a, b))
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
