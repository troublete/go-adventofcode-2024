package main

import (
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
)

type Map []Entry

type Entry struct {
	ID, N   int
	IsEmpty bool
}

type Layout []int

func (l Layout) LastNonEmpty() (int, int) {
	for i := len(l) - 1; i > 0; i-- {
		if l[i] != -1 {
			return i, l[i]
		}
	}
	return -1, -1
}

func (l Layout) ToString() string {
	sb := strings.Builder{}
	for _, e := range l {
		if e == -1 {
			sb.WriteString(".")
		} else {
			sb.WriteString(fmt.Sprintf("%d", e))
		}
	}
	return sb.String()
}

func (m Map) Layout() *Layout {
	l := Layout{}
	for _, e := range m {
		v := e.ID
		if e.IsEmpty {
			v = -1
		}
		for i := 0; i < e.N; i++ {
			l = append(l, v)
		}
	}
	return &l
}

func (l Layout) Improve(logger io.Writer) Layout {
	ret := l
	for id, e := range ret {
		if e == -1 {
			pos, v := ret.LastNonEmpty()
			if pos > id {
				ret[id] = v
				ret[pos] = -1
				if logger != nil {
					_, _ = logger.Write([]byte(ret.ToString() + "\n"))
				}
			}
		}
	}
	return ret
}

func (l Layout) Checksum() *big.Int {
	sum := big.NewInt(0)
	for i, v := range l {
		sum.Add(sum, big.NewInt(int64(i*v)))
	}
	return sum
}

func MapFromInput(input []byte) Map {
	m := Map{}

	id := 0
	newEntry := func(n int64, empty bool) Entry {
		e := Entry{
			ID:      -1,
			N:       int(n),
			IsEmpty: empty,
		}

		if !empty {
			e.ID = id
			id += 1
		}

		return e
	}

	empty := false
	for _, b := range input {
		i, _ := strconv.ParseInt(string(b), 10, 64)
		m = append(m, newEntry(i, empty))
		empty = !empty
	}

	return m
}
