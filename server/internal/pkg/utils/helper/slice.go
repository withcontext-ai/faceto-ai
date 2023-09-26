package helper

func SliceStringEqualUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[string]int)

	for _, value := range a {
		m[value]++
	}

	for _, value := range b {
		if m[value] > 0 {
			m[value]--
		} else {
			return false
		}
	}

	return true
}

func SliceIntEqualUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[int]int)

	for _, value := range a {
		m[value]++
	}

	for _, value := range b {
		if m[value] > 0 {
			m[value]--
		} else {
			return false
		}
	}

	return true
}

func RemoveDuplicates(s []string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0, len(s))

	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

func Intersect(a []string, b []string) []string {
	resp := make([]string, 0, len(b))
	mp := make(map[string]bool)

	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			resp = append(resp, s)
		}
	}

	return resp
}

func IndexOfStrSlice(s []string, a string) (bool, int) {
	for k, v := range s {
		if v == a {
			return true, k
		}
	}
	return false, 0
}
