package engine

import "sort"
// func GetRand(min int, max int) int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn((max-min)+1) + min
// }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func minMaxIntSlice (v []int) (int, int) {
	sort.Ints(v)
	return v[0], v[len(v)-1]
}