package utils

import (
	"fmt"
	"testing"
)

func TestIsPrime_97(t *testing.T) {
	// arrange
	no := 97
	expectedResult := true

	// act
	actualResult := IsPrime(no)

	// assert
	if actualResult != expectedResult {
		t.Errorf("IsPrime(97), expected = %v but actual = %v\n", expectedResult, actualResult)
	}
}

func TestIsPrime(t *testing.T) {
	var testData = []struct {
		no       int
		expected bool
	}{
		{no: 79, expected: true},
		{no: 91, expected: false},
		{no: 93, expected: false},
		{no: 97, expected: true},
	}
	for _, td := range testData {
		t.Run(fmt.Sprintf("TestIsPrime[%d]", td.no), func(t *testing.T) {

			// act
			actualResult := IsPrime(td.no)

			// assert
			if actualResult != td.expected {
				t.Errorf("IsPrime(%d), expected = %v but actual = %v\n", td.no, td.expected, actualResult)
			}
		})
	}
}
