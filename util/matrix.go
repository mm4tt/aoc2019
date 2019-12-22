package util

import (
	"fmt"
	"github.com/go-errors/errors"
	"strings"
)

type Matrix struct {
	NumRows, NumCols int
	M                [][]int
}

func NewMatrix(NumRows, NumCols int) Matrix {
	m := Matrix{
		NumRows: NumRows,
		NumCols: NumCols,
		M:       make([][]int, NumRows),
	}
	for i := range m.M {
		m.M[i] = make([]int, NumCols)
	}
	return m
}

func (m Matrix) String() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("Rows: %d, Cols: %d", m.NumRows, m.NumCols))
	for i := range m.M {
		b.WriteString("\n")
		b.WriteString(fmt.Sprint(m.M[i]))
	}
	return b.String()
}

func (m Matrix) Multiply(m1 Matrix, modulo int) (Matrix, error) {
	if m.NumCols != m1.NumRows {
		return Matrix{}, errors.Errorf("cannot multiple matrices, wrong sizes (%d, %d) * (%d, %d)",
			m.NumRows, m.NumCols, m1.NumRows, m1.NumCols)
	}
	ret := NewMatrix(m.NumRows, m1.NumCols)
	for r := range ret.M {
		for c := range ret.M[r] {
			for k := 0; k < m.NumCols; k++ {
				ret.M[r][c] = (ret.M[r][c] + SafeMultModulo(m.M[r][k], m1.M[k][c], modulo)) % modulo
			}
		}
	}
	return ret, nil
}
