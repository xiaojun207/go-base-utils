package utils

import (
	"fmt"
	"github.com/hpcloud/tail"
	"strconv"
	"testing"
	"time"
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

func TestReadFileLast(t *testing.T) {
	readChannel := make(chan string)
	filename := "test" + ".csv"
	go ReadFileLast(filename, readChannel)
	for r := range readChannel {
		fmt.Println(r)
	}
}

func TestReadFromCsv2(t *testing.T) {
	filename := "test" + ".csv"
	go read(filename)
	filename2 := "test2" + ".csv"
	read(filename2)
}

func read(filename string) {
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		// Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println(filename, " msg:", msg.Text)
	}
}
