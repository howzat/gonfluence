package filtering

import "fmt"

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func Matching(vs []string, predicate func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if predicate(v) {
			fmt.Printf("included element Matching predicate [%q]\n", v)
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func NotMatching(vs []string, predicate func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if !predicate(v) {
			fmt.Printf("included element NotMatching predicate [%q]\n", v)
			vsf = append(vsf, v)
		}
	}
	return vsf
}

