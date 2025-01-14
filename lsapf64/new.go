package lsapf64

import "math"

// New linear sum assignment problem (LSAP) from a square cost matrix.
// Note: mutates the cost matrix.
func New(A [][]float64) *LSAP {
	n := len(A)
	if n < 1 {
		panic("empty cost matrix")
	}
	for _, row := range A {
		if len(row) != n {
			panic("cost matrix not square")
		}
	}

	f := make([]int, n)
	fBar := make([]int, n)
	p := make([]int, n)
	c := make([]int, n)
	pi := make([]float64, n)

	for i := 0; i < n; i++ {
		f[i] = -1
		fBar[i] = -1
		p[i] = -1
		c[i] = -1
		pi[i] = -1
	}

	a := &LSAP{
		M:    math.Pow(1000, 3),
		a:    A,
		u:    make([]float64, n),
		v:    make([]float64, n),
		f:    f,
		fBar: fBar,
		p:    p,
		c:    c,
		pi:   pi,
		n:    n,
	}
	a.initialize()
	return a
}
