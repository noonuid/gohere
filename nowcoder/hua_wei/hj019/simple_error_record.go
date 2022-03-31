package hj019

import (
	"fmt"
	"strings"
)

// 错误记录的数据结构
type Record struct {
	filename   string
	lineNumber int
	count      int
}

func newRecord(path string, lineNumber int) *Record {
	var filename = getFilename(path)
	return &Record{filename: filename, lineNumber: lineNumber, count: 1}
}

// 根据路径设置文件名
func getFilename(path string) string {
	var lastBackslashIndex = strings.LastIndex(path, "\\")
	var suffix = path[lastBackslashIndex+1:]
	if len(suffix) > 16 {
		suffix = suffix[len(suffix)-16:]
	}
	return suffix
}

// 判断两个错误记录是否相同
func (record *Record) isSameAs(target *Record) bool {
	return record.filename == target.filename &&
		record.lineNumber == target.lineNumber
}

func (record *Record) print() {
	fmt.Printf("%s %d %d\n", record.filename, record.lineNumber, record.count)
}

type Table struct {
	data []*Record
}

func NewTable() *Table {
	return &Table{data: make([]*Record, 0)}
}

func (table *Table) insert(record *Record) bool {
	var index = table.findIndex(record)
	if index == -1 {
		table.data = append(table.data, record)
	} else {
		table.data[index].count++
	}
	return true
}

// 查找错误记录是否存在
func (table *Table) findIndex(record *Record) int {
	for i := 0; i < len(table.data); i++ {
		if table.data[i].isSameAs(record) {
			return i
		}
	}
	return -1
}

func Do() {
	var table = NewTable()

	var path string
	var lineNumber int
	var scanN1, _ = fmt.Scan(&path)
	var scanN2, _ = fmt.Scan(&lineNumber)
	for scanN1 > 0 && scanN2 > 0 {
		table.insert(newRecord(path, lineNumber))

		scanN1, _ = fmt.Scan(&path)
		scanN2, _ = fmt.Scan(&lineNumber)
	}

	if len(table.data) > 8 {
		table.data = table.data[len(table.data)-8:]
	}

	for _, record := range table.data {
		record.print()
	}
}
