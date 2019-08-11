package ap

// A Permutation P maps elements from one set U = {0,...,n-1} to another set
// V = {0,...,n-1}, where P[u] = v means that u ∈ U is assigned to v ∈ V.
//
//     a := ap.Permutation{1, 0, 2} // Assign 0 to 1, 1 to 0, and 2 to itself.
type Permutation []int

// Inverse converts an assignment from a set U to a set V to an assignment from
// V to U. If, for example, U is the left hand side of a bipartite matching and
// V is the right hand side, this function essentially swaps their sides.
//
//     p := ap.Permutation{1, 3, 2, 0}
//     p.Inverse() // {3, 0, 2, 1}
func (p Permutation) Inverse() Permutation {
	p2 := make(Permutation, len(p))
	for u, v := range p {
		p2[v] = u
	}
	return p2
}

// Matrix converts a permutation into a square matrix.
func (p Permutation) Matrix() Matrix {
	m := make(Matrix, len(p))
	for u, v := range p {
		m[u] = make([]bool, len(p))
		m[u][v] = true
	}
	return m
}