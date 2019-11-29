package main

import (
	"fmt"
)

// "strconv"

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	mot := append(a, b...)
	fmt.Println("mot: ", mot)

	// COPY
	// c1
	hai := make([]string, len(mot))
	copy(hai, mot)
	// c2
	hai = append([]string(nil), mot...)
	fmt.Println("hai: ", hai)

	// CUT
	fmt.Println(hai[2:], hai[:2])

	// DELETE
	// idx = 5,6,7
	ba := append(hai[:5], hai[8:]...)
	fmt.Println("Ba: ", ba, ba[2:])

	// c := ba[:2+copy(ba[2:], ba[7:])]
	d := append(ba[:2], ba[7:]...)
	fmt.Println("c: ", d, ba)

	//
	ba[2] = ba[5]
	fmt.Println(ba[(len(ba) - 4):])

	e := append(make([]string, 4), ba[(len(ba)-4):]...)
	fmt.Println("e: ", e)

	values := []string{"ABC", "CAB", "BCA"}
	factor := []int{100, 10, 1}

	for v := range values {
		hashKey := 0
		f := 0
		bytes := []byte(values[v])
		for i := range bytes {
			hashKey += int(bytes[i]) * factor[f]
			f++
		}
		fmt.Println(hashKey)
	}
}
