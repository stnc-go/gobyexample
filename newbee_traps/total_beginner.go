package newbee_traps

import "fmt"

func NilInitVariableWithExplicitType() interface{} {
	var x interface{} = nil
	return x
}

func NilInitSlicesAndMaps() (map[string]int, []int) {
	m := make(map[string]int)
	m["one"] = 1

	var s []int
	s = append(s, 1)

	return m, s
}

func InitStrings() string {
	var s string

	if s == "" {
		s = "default"
	}

	return s
}

func RangeSlices() {
	x := []string{"a", "b", "c"}

	for _, v := range x {
		fmt.Println(v)
	}
}

func MultiDimension() [][]int {
	x := 2
	y := 4

	table := make([][]int, x)
	for i := range table {
		table[i] = make([]int, y)
	}

	fmt.Println(table)

	return table
}

func ImmutableStrings() (string, string) {
	x := "test"
	xbytes := []byte(x)
	xbytes[0] = 'T'
	y := "sj"
	yrunes := []rune(y)
	yrunes[0] = '世'
	yrunes[1] = '界'

	y = string(yrunes)

	fmt.Println(x[0])
	fmt.Println(y[0:3])

	return string(xbytes), y
}
