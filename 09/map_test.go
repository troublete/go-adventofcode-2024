package main

import (
	"fmt"
	"math/big"
	"os"
	"testing"
)
import _ "embed"

//go:embed example.txt
var example []byte

func Test_MapFromInput(t *testing.T) {
	m := MapFromInput(example)

	for id, e := range []Entry{
		{IsEmpty: false, N: 2, ID: 0},
		{IsEmpty: true, N: 3, ID: -1},
		{IsEmpty: false, N: 3, ID: 1},
		{IsEmpty: true, N: 3, ID: -1},
		{IsEmpty: false, N: 1, ID: 2},
		{IsEmpty: true, N: 3, ID: -1},
		{IsEmpty: false, N: 3, ID: 3},
		{IsEmpty: true, N: 1, ID: -1},
		{IsEmpty: false, N: 2, ID: 4},
		{IsEmpty: true, N: 1, ID: -1},
		{IsEmpty: false, N: 4, ID: 5},
		{IsEmpty: true, N: 1, ID: -1},
		{IsEmpty: false, N: 4, ID: 6},
		{IsEmpty: true, N: 1, ID: -1},
		{IsEmpty: false, N: 3, ID: 7},
		{IsEmpty: true, N: 1, ID: -1},
		{IsEmpty: false, N: 4, ID: 8},
		{IsEmpty: true, N: 0, ID: -1},
		{IsEmpty: false, N: 2, ID: 9},
	} {
		if e.N != m[id].N || e.ID != m[id].ID || e.IsEmpty != m[id].IsEmpty {
			t.Errorf("#%v wrong entry, expected %#v, got %#v", id, e, m[id])
		}
	}
}

func TestMapLayout(t *testing.T) {
	m := MapFromInput(example)
	l := m.Layout().ToString()
	if l != "00...111...2...333.44.5555.6666.777.888899" {
		t.Error("expected different layout")
	}
}

func TestMapImprove(t *testing.T) {
	m := MapFromInput(example)
	l := m.Layout().Improve(os.Stdout)
	if l.ToString() != "0099811188827773336446555566" {
		t.Error("expected different layout")
	}
}

func TestMapChecksum(t *testing.T) {
	m := MapFromInput(example)
	l := m.Layout().Improve(nil)
	sum := l.Checksum()
	if sum.Cmp(big.NewInt(1928)) != 0 {
		t.Error("expected different checksum")
	}
}

func TestSimpleExample(t *testing.T) {
	m := MapFromInput([]byte("12345"))
	fmt.Printf("%v\n", m.Layout().ToString())
	m.Layout().Improve(os.Stdout)

}
