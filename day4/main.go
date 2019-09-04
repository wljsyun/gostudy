package main

import "fmt"

type Students struct {
	Age  int
	Name string
}

func main() {
	/*
			bool
			byte
			int
			uint
			float
			complex64 128
			uintptr
			string
			array
			struct
			function
			interface


			map
			slice
			channel

		s := "abc"
		fmt.Println([]byte(s))
		stu := new(Students)
		fmt.Println(*stu)
		fmt.Println(math.MinInt8, math.MaxInt16)
		m := make(map[int]string)
		m[1] = "1"
		m[2] = "2"
		fmt.Println(m)

		s1 := make([]int, 0)
		s1 = append(s1, 12)
		fmt.Println(s1)

		c := make(chan int,2)
		c <- 1
		fmt.Println(<-c)
	*/
	a := 123
	b := byte(a)
	fmt.Println(b)
}
