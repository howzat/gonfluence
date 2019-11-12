package filtering

import "fmt"

type StringPredicate = func(string) bool

func Any(vs []string, fn StringPredicate) bool {
	for _, v := range vs {
		if fn(v) {
			return true
		}
	}
	return false
}

func Matching(vs []string, fn StringPredicate) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if fn(v) {
			fmt.Printf("included element Matching StringPredicate [%q]\n", v)
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func NotMatching(vs []string, fn StringPredicate) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if !fn(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func FirstMatching(vs []string, fn StringPredicate) string {
	for _, v := range vs {
		if fn(v) {
			return v
		}
	}
	return ""
}

func Distinct(vs []string) []string {

	var instances = make(map[string]string)
	var results []string

	for _, key := range vs {
		_, ok := instances[key]
		if !ok {
			instances[key] = key
			results = append(results, key)
		}
	}

	return results
}
