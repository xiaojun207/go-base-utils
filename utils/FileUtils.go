package utils

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func SaveToFile(filename string, context string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666) //创建文件
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
		panic(err)
	}
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
	SaveToFile(filename, buffer.String()) //创建文件
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
	SaveToFile(filename, buffer.String()) //创建文件
}

func ReadFromCsv(filename string) [][]string {
	f, err := os.Open(filename) //打开文件
	if err != nil {
		log.Println(err)
		return [][]string{}
	}
	r := csv.NewReader(f)
	for {
		recordAll, err := r.ReadAll()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		return recordAll
	}
	return [][]string{}
}

func ReadFileLast(filename string, readChannel chan string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err.Error()))
	}
	//移动到文件末尾
	f.Seek(0, io.SeekEnd)
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadBytes('\n')
		//fmt.Println(err)
		if err == io.EOF {
			time.Sleep(time.Second)
			continue
		} else if err != nil {
			panic("ReadBytes error:" + err.Error())
		}

		lineStr := strings.TrimSpace(string(line))
		readChannel <- lineStr
	}
}
