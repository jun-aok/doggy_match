package util

import "fmt"

func Where[T comparable](a []T, f func(s T) bool) []T {
	ans := make([]T, 0)
	for _, x := range a {
		if f(x) {
			ans = append(ans, x)
		}
	}
	return ans
}

func Select[T any, TResult any](a []T, f func(s T) TResult) []TResult {
	ans := make([]TResult, 0)
	for _, x := range a {
		ans = append(ans, f(x))
	}
	return ans
}

func ToDictionary[T any, TKey comparable](a []T, f func(k T) TKey) (map[TKey]T, error) {
	m := make(map[TKey]T)
	for _, x := range a {
		id := f(x)
		if _, ok := m[id]; ok {
			return nil, fmt.Errorf("重複")
		}
		m[id] = x
	}
	return m, nil
}
