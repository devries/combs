package combs

import (
	"testing"
)

func TestCombinations(t *testing.T) {
	a := []int{1, 2, 3}

	combs := [][]int{
		[]int{1, 2},
		[]int{1, 3},
		[]int{2, 3},
	}

	for r := range Combinations(2, a) {
		for i, r2 := range combs {
			d := make(map[int]bool)
			for _, r3 := range r {
				d[r3] = true
			}

			if d[r2[0]] && d[r2[1]] {
				combs[i] = combs[len(combs)-1]
				combs = combs[:len(combs)-1]
				break
			}
		}
	}

	if len(combs) > 0 {
		t.Errorf("Not all combinations of %v were found", a)
	}
}

func TestPermutations(t *testing.T) {
	a := []int{1, 2, 3}

	perms := [][]int{
		[]int{1, 2, 3},
		[]int{2, 1, 3},
		[]int{3, 1, 2},
		[]int{1, 3, 2},
		[]int{2, 3, 1},
		[]int{3, 2, 1},
	}

	for r := range Permutations(a) {
		for i, r2 := range perms {
			if r[0] == r2[0] && r[1] == r2[1] && r[2] == r2[2] {
				perms[i] = perms[len(perms)-1]
				perms[len(perms)-1] = nil
				perms = perms[:len(perms)-1]
				break
			}
		}
	}

	if len(perms) > 0 {
		t.Errorf("Not all permutations of %v were found", a)
	}
}
