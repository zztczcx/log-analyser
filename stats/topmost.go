package stats

import (
	"sort"
)

// if value is the same, sort by key alphabetically
func TopMost(m map[string]int, n int) []string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if m[keys[i]] == m[keys[j]] {
			return keys[i] < keys[j]
		} else {
			return m[keys[i]] > m[keys[j]]
		}
	})

	if n > len(keys) {
		return keys
	}
	return (keys)[:n]
}
