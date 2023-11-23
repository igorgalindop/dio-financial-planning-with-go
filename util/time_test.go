package util

import (
	"strconv"
	"testing"
)

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2023-10-11T20:25:00")

	if convertedTime.Year() != 2023 {
		t.Errorf("Ano esperado 2023 - Ano retornado " + strconv.Itoa(convertedTime.Year()))
	}

	if convertedTime.Month() != 10 {
		t.Errorf("Mês esperado 10 - Mês retornado " + convertedTime.Month().String())
	}

	if convertedTime.Day() != 11 {
		t.Errorf("Dia esperado 11 - Dia retornado " + strconv.Itoa(convertedTime.Day()))
	}
}
