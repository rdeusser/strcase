package strcase

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"unicode"
)

var replacements = sync.Map{}

// ConfigureReplacement allows you to add additional words which will be considered for replacement.
func ConfigureReplacement(key, val string, ignoreCase bool) {
	if !ignoreCase {
		replacements.Store(key, val)
		return
	}
	for _, s := range generateReplacements(key) {
		replacements.Store(s, val)
	}
}

// generateReplacements returns all upper and lower case permutations of a key.
func generateReplacements(key string) []string {
	p := permutations(len(key))
	v := make([]string, 0)
	for i := 0; i < len(p); i++ {
		var buf bytes.Buffer
		for j, r := range key {
			switch p[i][j] {
			case 0:
				buf.WriteRune(unicode.ToLower(r))
			case 1:
				buf.WriteRune(unicode.ToUpper(r))
			}
		}
		v = append(v, buf.String())
	}
	return v
}

// permutations is a sloppy attempt at generating all permutations of 0 and 1 with n length.
func permutations(n int) [][]int {
	k := numPermutations(n)
	p := make([][]int, 0)
	i := 0
	for i < k {
		v := generate(n)
		p = append(p, v)
		p = uniq(p)
		i = len(p)
	}
	return p
}

func numPermutations(n int) int {
	if n < 0 {
		panic("n is negative")
	}
	return int(math.Pow(2, float64(n)))
}

func generate(n int) []int {
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = rand.Intn(2)
	}
	return v
}

func uniq(n [][]int) [][]int {
	v := make([][]int, 0)
	seen := make(map[string]struct{})
	for _, i := range n {
		s := fmt.Sprint(i)
		if _, ok := seen[s]; !ok {
			v = append(v, i)
			seen[s] = struct{}{}
		}
	}
	return v
}
