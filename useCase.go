package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Agregation(data []ValCurs) map[string][]DateRate {
	result := make(map[string][]DateRate)
	for _, dat := range data {
		for _, da := range dat.Valute {
			raw := strings.Replace(da.Value, ",", ".", -1)
			v, err := strconv.ParseFloat(raw, 64)
			if err != nil {
				continue
			}
			v = v / float64(da.Nominal)
			entry := DateRate{Date: dat.Date, Rate: v}
			result[da.CharCode] = append(result[da.CharCode], entry)
		}
	}
	return result
}

func Min(data map[string][]DateRate) {
	for key, entries := range data {
		minEntry := entries[0]
		for _, e := range entries[1:] {
			if e.Rate < minEntry.Rate {
				minEntry = e
			}
		}
		fmt.Printf("Min %s = %.4f on %s\n", key, minEntry.Rate, minEntry.Date)
	}
}

func Max(data map[string][]DateRate) {
	for key, entries := range data {
		maxEntry := entries[0]
		for _, e := range entries[1:] {
			if e.Rate > maxEntry.Rate {
				maxEntry = e
			}
		}
		fmt.Printf("Max %s = %.4f on %s\n", key, maxEntry.Rate, maxEntry.Date)
	}
}

func AVG(data map[string][]DateRate) {
	for key, entries := range data {
		sum := 0.0
		for _, e := range entries {
			sum += e.Rate
		}
		avg := sum / float64(len(entries))
		fmt.Printf("Avg %s = %.4f\n", key, avg)
	}
}
