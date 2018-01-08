package main

import "fmt"

// this is a comment

func main() {
	//base
	fmt.Println("Hello World")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[0])
	fmt.Println("Helo " + "World")
	var x string = "Hello World"
	fmt.Println(x)

	y := "Hello World"
	fmt.Println(y)

	z := "test"
	fmt.Println(z)

	/*
		i := 1
		for i <= 10 {
			fmt.Println(i)
			i = i + 1
		}
	*/

	//for if switch
	for i := 1; i <= 100; i++ {
		// fmt.Println(i)
		if i%7 == 0 {
			fmt.Println("even")
		} else if i%100 == 0 {
			fmt.Println("odd 100")
		}

		switch i {
		case 0:
			fmt.Println("Zero")
		case 100:
			fmt.Println("one hundred")
		default:
			fmt.Println("Unkown Number")
		}
	}

	//arrays
	var n [5]int
	n[4] = 100
	fmt.Println(n)
	fmt.Println(n[0])

	// m := [5]float64{ 100,101,202,303,404 }
	var m [5]float64
	m[0] = 100
	m[1] = 101
	m[2] = 202
	m[3] = 303
	m[4] = 404

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += m[i]
	}
	fmt.Println(total / 5)

	//slice
	//var o []float64
	o := make([]float64, 5, 10)
	arr_00 := []float64{1, 2, 3, 4, 5}
	arr_01 := arr_00[0:1] //?
	arr_02 := arr_00[2:4]
	arr_03 := arr_00[3:5]
	fmt.Println(o, arr_00, arr_01, arr_02, arr_03)

	//slice append copy
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	//maps
	//var demo_map map[string]int
	demo_map := make(map[string]int)
	demo_map["keys"] = 100
	fmt.Println(demo_map["keys"])

	demo_map_01 := make(map[int]int)
	demo_map_01[1] = 1111
	fmt.Println(demo_map_01[1])

	delete(demo_map_01, 1)
	fmt.Println(demo_map_01[1])
	fmt.Println(demo_map_01[2]) //?

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	fmt.Println(elements["H"])

	name, ok := elements["Un"]
	fmt.Println(name, ok)

	name1, ok := elements["H"]
	fmt.Println(name1, ok)

	if name, ok := elements["H"]; ok {
		fmt.Println(name, ok)
	}

	elements1 := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
	}

	name2, ok := elements1["He"]
	fmt.Println(name2, ok)

	name3, ok := elements1["L"]
	fmt.Println(name3, ok)
}
