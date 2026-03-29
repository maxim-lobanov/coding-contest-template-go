package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

type point struct{ r, c int }

func normalized(s []point) []point {
	minR, minC := s[0].r, s[0].c
	for _, p := range s {
		if p.r < minR {
			minR = p.r
		}
		if p.c < minC {
			minC = p.c
		}
	}
	out := make([]point, len(s))
	for i, p := range s {
		out[i] = point{p.r - minR, p.c - minC}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].r != out[j].r {
			return out[i].r < out[j].r
		}
		return out[i].c < out[j].c
	})
	return out
}

func orientations(shape []point) [][]point {
	seen := map[string]bool{}
	var result [][]point
	cur := append([]point{}, shape...)
	for flip := 0; flip < 2; flip++ {
		for range 4 {
			norm := normalized(cur)
			key := fmt.Sprintf("%v", norm)
			if !seen[key] {
				seen[key] = true
				result = append(result, norm)
			}
			next := make([]point, len(cur))
			for i, p := range cur {
				next[i] = point{p.c, -p.r}
			}
			cur = next
		}
		flipped := make([]point, len(shape))
		for i, p := range shape {
			flipped[i] = point{p.r, -p.c}
		}
		cur = flipped
	}
	return result
}

// DLX using parallel arrays
type dlx struct {
	L, R, U, D, C, S []int
	nPrimary         int
}

func newDLX(nPrimary, nSecondary int) *dlx {
	n := nPrimary + nSecondary + 1
	d := &dlx{
		L: make([]int, n), R: make([]int, n),
		U: make([]int, n), D: make([]int, n),
		C: make([]int, n), S: make([]int, n),
		nPrimary: nPrimary,
	}
	for i := range n {
		d.U[i], d.D[i], d.C[i] = i, i, i
	}
	if nPrimary > 0 {
		d.R[0], d.L[0] = 1, nPrimary
		d.L[1], d.R[nPrimary] = 0, 0
		for i := 1; i < nPrimary; i++ {
			d.R[i], d.L[i+1] = i+1, i
		}
	}
	for i := nPrimary + 1; i < n; i++ {
		d.L[i], d.R[i] = i, i
	}
	return d
}

func (d *dlx) addRow(cols []int) {
	first := len(d.L)
	for i, c := range cols {
		c++
		idx := len(d.L)
		d.C = append(d.C, c)
		d.U = append(d.U, d.U[c])
		d.D = append(d.D, c)
		d.D[d.U[c]] = idx
		d.U[c] = idx
		if i == 0 {
			d.L = append(d.L, idx)
			d.R = append(d.R, idx)
		} else {
			d.L = append(d.L, idx-1)
			d.R = append(d.R, first)
			d.R[idx-1] = idx
			d.L[first] = idx
		}
		d.S[c]++
	}
}

func (d *dlx) cover(c int) {
	d.R[d.L[c]], d.L[d.R[c]] = d.R[c], d.L[c]
	for i := d.D[c]; i != c; i = d.D[i] {
		for j := d.R[i]; j != i; j = d.R[j] {
			d.D[d.U[j]], d.U[d.D[j]] = d.D[j], d.U[j]
			d.S[d.C[j]]--
		}
	}
}

func (d *dlx) uncover(c int) {
	for i := d.U[c]; i != c; i = d.U[i] {
		for j := d.L[i]; j != i; j = d.L[j] {
			d.S[d.C[j]]++
			d.D[d.U[j]], d.U[d.D[j]] = j, j
		}
	}
	d.R[d.L[c]], d.L[d.R[c]] = c, c
}

func (d *dlx) solve() bool {
	if d.R[0] == 0 {
		return true
	}
	best, minS := -1, len(d.L)
	for j := d.R[0]; j != 0; j = d.R[j] {
		if d.S[j] < minS {
			minS, best = d.S[j], j
		}
	}
	if minS == 0 {
		return false
	}
	d.cover(best)
	for r := d.D[best]; r != best; r = d.D[r] {
		for j := d.R[r]; j != r; j = d.R[j] {
			d.cover(d.C[j])
		}
		if d.solve() {
			return true
		}
		for j := d.L[r]; j != r; j = d.L[j] {
			d.uncover(d.C[j])
		}
	}
	d.uncover(best)
	return false
}

func canFit(W, H int, allOrients [][][]point, counts []int) bool {
	total := 0
	for _, c := range counts {
		total += c
	}
	if total == 0 {
		return true
	}
	if total*len(allOrients[0][0]) > W*H {
		return false
	}

	d := newDLX(total, W*H)
	pieceIdx := 0
	for sid, cnt := range counts {
		for range cnt {
			for _, orient := range allOrients[sid] {
				maxR, maxC := 0, 0
				for _, p := range orient {
					if p.r > maxR {
						maxR = p.r
					}
					if p.c > maxC {
						maxC = p.c
					}
				}
				for r := 0; r+maxR < H; r++ {
					for c := 0; c+maxC < W; c++ {
						cols := []int{pieceIdx}
						for _, p := range orient {
							cols = append(cols, total+(r+p.r)*W+(c+p.c))
						}
						d.addRow(cols)
					}
				}
			}
			pieceIdx++
		}
	}
	return d.solve()
}

func parseShapes(input []string) ([][]point, int) {
	var shapes [][]point
	i := 0
	for i < len(input) {
		if strings.Contains(input[i], "x") {
			break
		}
		if strings.HasSuffix(input[i], ":") {
			var s []point
			i++
			for row := 0; i < len(input) && len(input[i]) > 0 &&
				!strings.HasSuffix(input[i], ":") && !strings.Contains(input[i], "x"); row++ {
				for c, ch := range input[i] {
					if ch == '#' {
						s = append(s, point{row, c})
					}
				}
				i++
			}
			shapes = append(shapes, s)
		} else {
			i++
		}
	}
	return shapes, i
}

func solution(input []string) string {
	shapes, regionStart := parseShapes(input)
	allOrients := make([][][]point, len(shapes))
	for i, s := range shapes {
		allOrients[i] = orientations(s)
	}

	result := 0
	for i := regionStart; i < len(input); i++ {
		line := input[i]
		if !strings.Contains(line, "x") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		dims := strings.Split(strings.TrimSpace(parts[0]), "x")
		W, H := cast.ParseInt(dims[0]), cast.ParseInt(dims[1])
		counts := cast.ParseIntArray(strings.TrimSpace(parts[1]))
		if canFit(W, H, allOrients, counts) {
			result++
		}
	}
	return cast.ToString(result)
}
