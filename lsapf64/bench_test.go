package lsapf64_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/xiaoxfan/ap/lsapf64"
)

func randomLSAP(n int) *lsapf64.LSAP {
	rand.Seed(time.Now().UTC().UnixNano())
	A := make([][]float64, n)
	for i := 0; i < n; i++ {
		Ai := make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				Ai[j] = float64(rand.Int63n(100000))
			}
		}
		A[i] = Ai
	}
	ap := lsapf64.New(A)
	return ap
}

func benchSolv(b *testing.B, size int, removes int) {
	for n := 0; n < b.N; n++ {
		ap := randomLSAP(size)
		assign := ap.Assign()

		// Incrementally remove assignments and re-optimize.
		for i := 0; i < removes; i++ {
			u := rand.Intn(size)
			v := assign[u]
			ap.Remove(u, v)
			assign = ap.Assign()
		}
	}
}

func benchCopy(b *testing.B, size int, copies int) {
	for n := 0; n < b.N; n++ {
		ap := randomLSAP(size)
		for i := 0; i < copies; i++ {
			ap.Copy()
		}
	}
}

func BenchmarkSolve10(b *testing.B)  { benchSolv(b, 10, 0) }
func BenchmarkSolve100(b *testing.B) { benchSolv(b, 100, 0) }
func BenchmarkSolve1K(b *testing.B)  { benchSolv(b, 1000, 0) }
func BenchmarkSolve10K(b *testing.B) { benchSolv(b, 10000, 0) }

func BenchmarkIncremental10x1(b *testing.B)   { benchSolv(b, 10, 1) }
func BenchmarkIncremental100x10(b *testing.B) { benchSolv(b, 100, 10) }
func BenchmarkIncremental1Kx100(b *testing.B) { benchSolv(b, 1000, 100) }

func BenchmarkCopy10x1K(b *testing.B)  { benchCopy(b, 10, 1000) }
func BenchmarkCopy100x1K(b *testing.B) { benchCopy(b, 100, 1000) }
func BenchmarkCopy1Kx1K(b *testing.B)  { benchCopy(b, 1000, 1000) }
