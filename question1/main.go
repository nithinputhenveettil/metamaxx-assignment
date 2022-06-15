package main

import "fmt"

func main() {
	c1 := []string{"Red", "Black", "White"}
	c2 := []string{"Black", "Yellow", "Orange"}

	m := make(map[string]int)

	addElementsToMap(c1, m)
	addElementsToMap(c2, m)

	l := len(m)
	c3 := make([]string, l)

	c := 0
	for i := range m {
		c3[c] = i
		c++
	}

	fmt.Println(c3)

}

func addElementsToMap(s []string, m map[string]int) {
	l := len(s)
	for i := 0; i < l; i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = i
		}
	}
}
