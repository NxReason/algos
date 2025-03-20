package list

import (
	"errors"
	"strconv"
)

type IntegerCell struct {
	Value int
	Next  *IntegerCell
}

type IntegerList struct {
	first   *IntegerCell
	last    *IntegerCell
	current *IntegerCell
}

func (il *IntegerList) Append(value int) {
	cell := IntegerCell{Value: value, Next: nil}
	if il.last != nil {
		il.last.Next = &cell
	} else {
		il.first = &cell
	}
	il.last = &cell
}

func (il *IntegerList) Prepend(value int) {
	cell := IntegerCell {Value: value, Next: il.first}
	il.first = &cell
	if il.last == nil {
		il.last = &cell
	}
}

func (il *IntegerList) Contains(value int) bool {
	il.current = il.first
	for il.current.Next != nil {
		if il.current.Value == value { return true }
		il.current = il.current.Next
	}
	return false
}

func (il *IntegerList) InsertAt(value int, pos int) error {
	if pos == 0 {
		il.Prepend(value)
		return nil
	}

	i := 1
	il.current = il.first
	for i < pos {
		i++
		if il.current.Next == nil { return errors.New("position is out of bounds of the list") }
		il.current = il.current.Next
	}
	cell := IntegerCell { Value: value, Next: il.current.Next }
	il.current.Next = &cell

	return nil
}

func (il *IntegerList) ToString() string {
	il.current = il.first
	result := ""
	for il.current.Next != nil {
		result += strconv.Itoa(il.current.Value) + " "
		il.current = il.current.Next
	}
	result += strconv.Itoa(il.current.Value)
	return result
}
