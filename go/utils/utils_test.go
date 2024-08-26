package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type ParseLineTestData struct {
	input   string
	city    string
	reading float64
}

func getParseLineTestData() []ParseLineTestData {
	return []ParseLineTestData{
		{input: "city;12.3", city: "city", reading: float64(12.3)},
		{input: "Murmansk;-0.2", city: "Murmansk", reading: float64(-0.2)},
		{input: "Cairns;24.0", city: "Cairns", reading: float64(24.0)},
		{input: "Medan;19.0", city: "Medan", reading: float64(19.0)},
		{input: "Yerevan;15.9", city: "Yerevan", reading: float64(15.9)},
		{input: "Flores,  Petén;30.1", city: "Flores,  Petén", reading: float64(30.1)},
		{input: "Ghanzi;27.6", city: "Ghanzi", reading: float64(27.6)},
		{input: "Charlotte;3.9", city: "Charlotte", reading: float64(3.9)},
		{input: "Albuquerque;11.6", city: "Albuquerque", reading: float64(11.6)},
		{input: "Los Angeles;21.5", city: "Los Angeles", reading: float64(21.5)},
		{input: "Ifrane;11.0", city: "Ifrane", reading: float64(11.0)},
	}
}

func TestParseLine(t *testing.T) {
	testData := getParseLineTestData()

	for _, data := range testData {
		city, reading := ParseLine(data.input)

		require.Equal(t, city, data.city, "City not parsed correctly")
		require.Equal(t, reading, data.reading, "Reading not parsed correctly")
	}

}

func TestParseLineRegex(t *testing.T) {
	testData := getParseLineTestData()

	for _, data := range testData {
		city, reading := ParseLineRegex(data.input)

		require.Equal(t, city, data.city, "City not parsed correctly")
		require.Equal(t, reading, data.reading, "Reading not parsed correctly")
	}

}

func BenchmarkParseLine(b *testing.B) {
	testData := getParseLineTestData()

	for i := 0; i < b.N; i++ {
		for _, data := range testData {
			ParseLine(data.input)
		}
	}
}

func BenchmarkParseLineRegex(b *testing.B) {
	testData := getParseLineTestData()

	for i := 0; i < b.N; i++ {
		for _, data := range testData {
			ParseLineRegex(data.input)
		}
	}
}
