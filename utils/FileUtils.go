package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func SaveToFile(filename string, context string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666) //创建文件
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
		panic(err)
	}
	defer f.Close()
	f.Seek(0, io.SeekEnd)

	f.WriteString(context)
}

func SaveToCsv(filename string, rows [][]string) {
	var buffer bytes.Buffer
	for _, row := range rows {
		s := strings.Join(row, ",")
		buffer.WriteString(s)
		buffer.WriteString("\r\n")
	}
	SaveToFile(filename+".csv", buffer.String()) //创建文件
}

func AppendToCsv(filename string, row []string) {
	SaveToCsv(filename, [][]string{row})
}

func AppendsToCsv(filename string, rows [][]string) {
	SaveToCsv(filename, rows)
}

func AppendFloatToCsv(filename string, row []float64) {
	AppendToCsv(filename, Float64sToStrings(row))
}

func AppendFloatsToCsv(filename string, rows [][]float64) {
	var buffer bytes.Buffer
	for _, row := range rows {
		s := strings.Join(Float64sToStrings(row), ",")
		buffer.WriteString(s)
		buffer.WriteString("\r\n")
	}
	SaveToFile(filename+".csv", buffer.String()) //创建文件
}
