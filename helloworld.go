package main

import "fmt"

const (
	a        = iota
	s string = "CONST"
)

func main() {
	//var str string
	i := 22
	for i := 0; i < 22; i++ {
		println("number is: ", i)
	}
	if i <= 20 {
		fmt.Printf("Variable i equel or lower than 20 \n")
	} else {
		fmt.Printf("Variable i bigger than 20 \n")
	}
	/*var (
		n int;
		b bool
	);*/
	/*	str = "Wappy New Year " +
				"My Dear Friend"
		c := []rune(str)
		c[0] = 'H'
		toStr := string(c)
		println(str + "\n")
		fmt.Printf("%v \n", i)
		//println(n)
		//println(b)
		println(s)
		fmt.Printf("%v \n", toStr)
	*/
	array := make(map[string]string)
	array["a"] = "lol"

	for index, value := range array {
		fmt.Println("index:", index, "value:", value)
	}

	ints := []int{1, 2, 3, 4, 5, 6, 7}

	//copy(ints[3:3], []int{9, 8, 0})

	fmt.Println(ints[3:3])

}
