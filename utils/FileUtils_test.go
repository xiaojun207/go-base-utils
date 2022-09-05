package utils

import (
	"strconv"
	"testing"
)

func TestSaveToFile(t *testing.T) {
	for i := 0; i < 10; i++ {
		temp := "\r\n" + strconv.Itoa(i) + ",wdfa"
		SaveToFile("text.txt", temp)
	}
}

func TestSaveToCsv(t *testing.T) {
	rows := [][]string{{"a5", "b5"}, {"a6", "b6"}, {"a7", "b7"}}
	SaveToCsv("test", rows)
	floatRows := [][]float64{{11, 11}, {22, 22}, {33, 33}}
	AppendFloatsToCsv("test", floatRows)
}

func TestReadFromCsv(t *testing.T) {
	filename := "test" + ".csv"
	ReadFromCsv(filename)
}
