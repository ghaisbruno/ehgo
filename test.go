package main

import (
	"fmt"
	"strconv"
)

func slice(size int) []string {
	s := make([]string, size) //slice (it has a variable size)
	return s
}

func main() {
	const N int = 10
	b := [N]string{"teste", "!"} //array
	fmt.Println(b)

	c := slice(N)
	sl := c[4:N]  //slicing a slice to get only a part of it
	sla := b[1:3] // slicing an array to get only a part of it

	for i := 0; i < N; i++ {
		b[i] = "ola"
		c[i] = b[i] + " mundo " + strconv.Itoa(i) + "!"
	}
	fmt.Println(c)
	fmt.Println(sl)
	fmt.Println(sla)
}
