package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
	s := `Month,Day
     January,Sunday
     February,Monday
     Web,Thusday`
	// readCsvFile(strings.NewReader(s))
	// readCsvFile2(strings.NewReader(s))

	cs := NewCSVScanner(strings.NewReader(s))
	for cs.Scan() {
		fmt.Printf("month: %s,day: %s\n", cs.Text("Month"), cs.Text("Day"))
		fmt.Println(cs.GetIndex())
	}
}

var (
	ErrColumnNotExists = errors.New("csv.Scanner: column not exists")
)

func readCsvFile(r io.Reader) error {
	reader := csv.NewReader(r)
	data, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for _, row := range data {
		fmt.Printf("month: %s,day: %s\n", row[0], row[1])
	}
	return nil
}

func readCsvFile2(r io.Reader) error {
	reader := csv.NewReader(r)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("month: %s,day: %s\n", row[0], row[1])
	}
	return nil
}

type CSVScanner struct {
	Reader    *csv.Reader
	Head      map[string]int
	Row       []string
	CurRowNum int
}

func NewCSVScanner(r io.Reader) CSVScanner {
	csvReader := csv.NewReader(r)
	head, err := csvReader.Read()
	if err != nil {
		return CSVScanner{}
	}
	h := make(map[string]int)
	for index, column := range head {
		h[column] = index
	}
	return CSVScanner{Reader: csvReader, Head: h, CurRowNum: 1}
}

func (cs *CSVScanner) Scan() bool {
	r, err := cs.Reader.Read()
	if err != nil {
		return false
	}
	cs.Row = r
	cs.CurRowNum += 1
	return true
}

func (cs *CSVScanner) Text(s string) string {
	if idx, ok := cs.Head[s]; ok {
		return cs.Row[idx]
	}
	return ""
}

func (cs *CSVScanner) TextOfIndex(index int) (string, error) {
	if index >= len(cs.Row) {
		return "", nil
	}
	return cs.Row[index], nil
}

func (cs *CSVScanner) GetIndex() int {
	return cs.CurRowNum
}
