package problem1071

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	str1 string
	str2 string
	ans  string
}{

	{
		"AAAAAAA",
		"AAAAAAAAAAA",
		"A",
	},

	{
		"NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM",
		"NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM",
		"NLZGM",
	},

	{
		"ABCABC",
		"ABC",
		"ABC",
	},

	{
		"ABABAB",
		"ABAB",
		"AB",
	},

	{
		"LEET",
		"CODE",
		"",
	},

	// 可以有多个 testcase
}

func Test_gcdOfStrings(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, gcdOfStrings(tc.str1, tc.str2), "输入:%v", tc)
	}
}

func Benchmark_gcdOfStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			gcdOfStrings(tc.str1, tc.str2)
		}
	}
}
